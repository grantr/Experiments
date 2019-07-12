package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	vision "cloud.google.com/go/vision/apiv1"
	cloudevents "github.com/cloudevents/sdk-go"
)

var (
	finalizeEventType = "com.google.storage.finalize"
)

func processEvent(event cloudevents.Event) {
	// TODO: Validate that file is an image

	if event.Context == nil {
		fmt.Printf("event.Context is nil. cloudevents.Event\n%s\n", event.String())
		return
	}
	if event.Context.GetType() != finalizeEventType {
		fmt.Printf("Invalid event type %s. Supported event type: %s\n", event.Context.GetType(), finalizeEventType)
		return
	}
	pubsubMsg := pubsub.Message{}
	if err := event.DataAs(&pubsubMsg); err != nil {
		fmt.Printf("Error extracting data from event. Error:%s\n", err.Error())
		return
	}
	bucketID, err := getBucketID(&pubsubMsg)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	objectID, err := getobjectID(&pubsubMsg)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	objectURI := getObjectURI(bucketID, objectID)
	// fmt.Printf("ObjectURI:%s\n", objectURI)

	text, err := extractText(objectURI)
	if err != nil {
		fmt.Printf("Error extracting extracting labels. Error:%s", err.Error())
		return
	}

	fmt.Printf("Object:%s\nText:%s", objectURI, text)

	if text != "" {
		if err := writeToGcs(bucketID, objectID+".txt", text); err != nil {
			fmt.Printf(err.Error())
			return
		}
	}
}
func getBucketID(pubsubMsg *pubsub.Message) (string, error) {
	return getAttribute("bucketId", pubsubMsg)
}
func getobjectID(pubsubMsg *pubsub.Message) (string, error) {
	return getAttribute("objectId", pubsubMsg)
}

func getAttribute(attribute string, pubsubMsg *pubsub.Message) (string, error) {
	if pubsubMsg == nil {
		return "", fmt.Errorf("pubsubMsg is nil")
	}
	attributeValue, ok := pubsubMsg.Attributes[attribute]
	if !ok {
		return "", fmt.Errorf("%s not found in Data.Attributes", attribute)
	}
	return attributeValue, nil
}

func getObjectURI(bucketID, objectID string) string {
	return fmt.Sprintf("gs://%s/%s", bucketID, objectID)
}

func writeToGcs(bucketID string, objectID string, text string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	wc := client.Bucket(bucketID).Object(objectID).NewWriter(ctx)
	fmt.Fprintf(wc, text)
	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}

func main() {

	// customEvent := cloudevents.NewEvent(cloudevents.VersionV02)

	// customEvent.SetType("com.google.storage.finalize")
	// customEvent.SetData(pubsub.Message{
	// 	Attributes: map[string]string{
	// 		"bucketId": "akashv-bucket",
	// 		"objectId": "images/text.jpg",
	// 	},
	// })
	// processEvent(customEvent)

	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatal("Failed to create client, ", err)
	}
	log.Fatal(c.StartReceiver(context.Background(), processEvent))

}

func extractText(objectURI string) (string, error) {
	ctx := context.Background()
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return "", fmt.Errorf("Failed to create client: %v", err)
	}
	defer client.Close()

	image := vision.NewImageFromURI(objectURI)
	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return "", err
	}

	if len(annotations) == 0 {
		return "", nil
	}
	return annotations[0].Description, nil
}
