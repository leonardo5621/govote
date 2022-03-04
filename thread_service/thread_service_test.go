package thread_service


import (
	"testing"
	"reflect"
	"gotest.tools/v3/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreationvalidation(t *testing.T) {
	newThread := ThreadCreationPayload{
		Title: "tester",
		OwnerUserId:  "t!#@#WRefdssd",
		Description: "Test",
	}
	ownerUserIdValidationErr := newThread.Validate()
	if ownerUserIdValidationErr == nil {
		t.Error("UserId format validation failed, non-hexadecimal character were allowed")
	}
	oid := primitive.NewObjectID()
	validThread := ThreadCreationPayload{
		Title: "tester",
		OwnerUserId:  oid.Hex(),
		Description: "Test",
	}
	validationErrName := validThread.Validate()
	if validationErrName != nil {
		t.Error("Validation failed")
	}
}

func TestCommentvalidation(t *testing.T) {
	threadId := primitive.NewObjectID().Hex()
	userId := primitive.NewObjectID().Hex()
	commentPayload := &CreateCommentPayload{
		Text: "Just a comment",
		AuthorUserId: "12323@weerer",
		ThreadId: threadId,
	}
	authorValidation := commentPayload.Validate()
	if authorValidation == nil {
		t.Error("authorId validation failed, non-hexadecimal characters were allowed")
	}
	validComment := &CreateCommentPayload{
		Text: "Just a comment",
		AuthorUserId: userId,
		ThreadId: threadId,
	}
	idValidationErr := validComment.Validate()
	if idValidationErr != nil {
		t.Errorf("Id validation failed, valid ID has not been recognized, %v", idValidationErr)
	}
}

func TestInits(t *testing.T) {
	userId := primitive.NewObjectID().Hex()
	invalidThread := &ThreadCreationPayload{
		Title: "tester",
		OwnerUserId:  "asoudhqwe87ewqi12",
		Description: "Test",
	}
	thread := &ThreadModel{}
	err := thread.Init(invalidThread)
	if err == nil {
		t.Errorf("Init failed, a nox-hexadecimal userId should not be allowed")
	}
	validThread := &ThreadCreationPayload{
		Title: "tester",
		OwnerUserId:  userId,
		Description: "Test",
	}
	validThreadErr := thread.Init(validThread)
	if validThreadErr != nil {
		t.Errorf("Init failed, valid userId was not allowed")
	}

	threadId := primitive.NewObjectID().Hex()
	invalidComment := &CreateCommentPayload{
		Text: "Just a comment",
		AuthorUserId: "WhaTever",
		ThreadId: threadId,
	}
	comment := &CommentModel{}
	invalidAuthorErr := comment.Init(invalidComment)
	if invalidAuthorErr == nil {
		t.Errorf("Init failed, a nox-hexadecimal AuthoruserId should not be allowed")
	}
	invalidComment.AuthorUserId = userId
	invalidComment.ThreadId = "Nonvalid245ytgd"
	invalidThreadErr := comment.Init(invalidComment)
	if invalidThreadErr == nil {
		t.Errorf("Init failed, a nox-hexadecimal ThreadId should not be allowed")
	}

	validComment := &CreateCommentPayload{
		Text: "Just a comment",
		AuthorUserId: userId,
		ThreadId: threadId,
	}
	initValidationErr := comment.Init(validComment)
	if initValidationErr != nil {
		t.Errorf("Init comment failed with the right parameters")
	}
}

func TestThreadfinder(t *testing.T) {
	oid := primitive.NewObjectID()
	threadFinder := &ThreadFinder{}
	err := threadFinder.Init(oid.Hex())
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, reflect.DeepEqual(threadFinder.ThreadId, &oid), true)	
	assert.Equal(t, reflect.DeepEqual(*threadFinder.GetFindOneQuery(), bson.M{"_id": threadFinder.ThreadId }), true)
	assert.Equal(
		t,
		reflect.TypeOf(threadFinder.GetDecodeTargetStruct()).Kind(),
		reflect.TypeOf(&ThreadModel{}).Kind(),
	)
}
