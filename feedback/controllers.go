package feedback

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/slawek87/JobHunters/conf"
	"github.com/astaxie/beego/validation"
	"encoding/json"
	"errors"
)

const MongoDBIndex = "Feedback"

type FeedbackController struct {
	Feedback Feedback
}

type MessageController struct {
	Message Message
}

type MainController struct {
	FeedbackController FeedbackController
	MessageController  MessageController
}

func (controller *FeedbackController) SetFeedback(feedback Feedback) {
	controller.Feedback = feedback
}

func (controller *FeedbackController) SetFeedbackID(feedbackID string) {
	if feedbackID != "" {
		controller.Feedback.FeedbackID = bson.ObjectIdHex(feedbackID)
	}
}

func (controller *FeedbackController) SetCandidateID(candidateID string) {
	controller.Feedback.CandidateID = bson.ObjectIdHex(candidateID)
}

func (controller *FeedbackController) SetOfferID(offerID string) {
	controller.Feedback.OfferID = bson.ObjectIdHex(offerID)
}

func (controller *FeedbackController) GetFeedback() Feedback {
	return controller.Feedback
}

func (controller *MessageController) SetMessage(message Message) {
	controller.Message = message
}

func (controller *MessageController) SetMessageID(messageID string) {
	controller.Message.MessageID = bson.ObjectIdHex(messageID)
}

func (controller *MessageController) SetSenderID(senderID string) {
	controller.Message.SenderID = senderID
}

func (controller *MessageController) SetSenderFullName(senderFullName string) {
	controller.Message.SenderFullName = senderFullName
}

func (controller *MessageController) SetContent(content string) {
	controller.Message.Content = content
}

func (controller *MessageController) GetMessage() Message {
	return controller.Message
}

func (controller *MainController) SetRead(readerID string) {
	update := false

	for i, message := range controller.FeedbackController.Feedback.Messages {
		if message.IsRead == false && message.SenderID != readerID {
			controller.FeedbackController.Feedback.Messages[i].IsRead = true
			controller.FeedbackController.Feedback.Messages[i].UpdatedAt = time.Now()
			update = true
		}
	}

	if update == true {
		session, db := conf.MongoDB()
		defer session.Close()

		collection := db.C(MongoDBIndex)
		collection.Update(
			bson.M{"feedback_id": &controller.FeedbackController.Feedback.FeedbackID},
			&controller.FeedbackController.Feedback)
	}
}

func (controller *MainController) CreateFeedback() error {
	session, db := conf.MongoDB()
	defer session.Close()

	controller.FeedbackController.Feedback.FeedbackID = bson.NewObjectId()
	controller.FeedbackController.Feedback.CreatedAt = time.Now()
	controller.FeedbackController.Feedback.UpdatedAt = time.Now()

	valid := validation.Validation{}

	isValid, _ := valid.Valid(controller.FeedbackController.Feedback)

	if !isValid {
		errorMsg := make(map[string]string)
		for _, err := range valid.Errors {
			errorMsg[err.Field] = err.Message
		}
		results, _ := json.Marshal(errorMsg)
		return errors.New(string(results))
	}

	collection := db.C(MongoDBIndex)
	return collection.Insert(controller.FeedbackController.Feedback)
}

func (controller *MainController) CreateMessage() error {
	session, db := conf.MongoDB()
	defer session.Close()

	controller.MessageController.Message.MessageID = bson.NewObjectId()
	controller.MessageController.Message.CreatedAt = time.Now()
	controller.MessageController.Message.UpdatedAt = time.Now()

	collection := db.C(MongoDBIndex)

	return collection.Update(
		bson.M{"feedback_id": controller.FeedbackController.Feedback.FeedbackID},
		bson.M{"$push": bson.M{"messages": controller.MessageController.Message}})
}

func (controller *MainController) Get() error {
	session, db := conf.MongoDB()
	defer session.Close()

	return db.C(MongoDBIndex).Find(bson.M{
		    "feedback_id": controller.FeedbackController.Feedback.FeedbackID,
			"offer_id": controller.FeedbackController.Feedback.OfferID,
			"candidate_id": controller.FeedbackController.Feedback.CandidateID}).
				One(&controller.FeedbackController.Feedback)
}