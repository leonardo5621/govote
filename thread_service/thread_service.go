package thread_service

import (
	"context"
	"github.com/leonardo5621/govote/orm"
	"github.com/leonardo5621/govote/utilities"
	"github.com/leonardo5621/govote/user_service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ThreadServer struct {
	UnimplementedThreadServiceServer
}

func (s *ThreadServer) GetThread(ctx context.Context, req *GetThreadRequest) (*GetThreadResponse, error) {
	// Init the threadFinder, which has been set
	// To search a thread by its Id
	threadFinder := &ThreadFinder{}
	err := threadFinder.Init(req.GetThreadId())
	if err != nil {
		return nil, utilities.ReturnValidationError(err)
	}
	// Perform the search
	collection := orm.OrmSession.Client.Database("upvote").Collection("thread")
	thread, err := orm.FindDocument(threadFinder, collection, ctx)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	// Convert to the struct which corresponds to
	// The response model
	responseModel, err := orm.ConvertStruct(thread, Thread{})
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	return &GetThreadResponse{Thread: responseModel.(*Thread)}, nil
}

func (s *ThreadServer) CreateThread(ctx context.Context, req *CreateThreadRequest) (*CreateThreadResponse, error) {
	threadPayload := req.GetThread()
	// Message parameters validation
	validationError := threadPayload.ValidateAll()
	if validationError != nil {
		return nil, utilities.ReturnValidationError(validationError)
	}
	// Pass the message parameter to a ThreadModel struct, which
	// Has bson Marshal support
	thread, err := orm.ConvertStruct(threadPayload, ThreadModel{})
	assertedThread := thread.(*ThreadModel)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	// Check if user exists
	// And get its name if it is the case
	userFinder := &user_service.UserFinder{}
	errInvalidId := userFinder.Init(threadPayload.GetOwnerUserId())
	if errInvalidId != nil {
		return nil, utilities.ReturnValidationError(errInvalidId)
	}
	userCollection := orm.OrmSession.Client.Database("upvote").Collection("user")
	foundUser, err := orm.FindDocument(userFinder, userCollection, ctx)
	user := foundUser.(*user_service.UserModel)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	assertedThread.OwnerUserName = user.UserName
	// Thread creation
	collection := orm.OrmSession.Client.Database("upvote").Collection("thread")
	threadId, err := orm.Create(assertedThread, collection, ctx)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	return &CreateThreadResponse{ThreadId: threadId}, nil
}


func (s *ThreadServer) CreateComment(ctx context.Context, req *CreateCommentRequest) (*CreateCommentResponse, error) {
	db := orm.OrmSession.Client.Database("upvote")
	commentPayload := req.GetComment()
	// Message parameters validation
	validationError := commentPayload.ValidateAll()
	if validationError != nil {
		return nil, utilities.ReturnValidationError(validationError)
	}
	// Pass the message parameter to a ThreadModel struct, which
	// Has bson Marshal support
	comment, err := orm.ConvertStruct(commentPayload, CommentModel{})
	assertedComment := comment.(*CommentModel)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	// Check if user exists
	// And get its name if it is the case
	userFinder := &user_service.UserFinder{}
	errInvalidId := userFinder.Init(commentPayload.GetAuthorUserId())
	if errInvalidId != nil {
		return nil, utilities.ReturnValidationError(errInvalidId)
	}
	userCollection := db.Collection("user")
	foundUser, err := orm.FindDocument(userFinder, userCollection, ctx)
	user := foundUser.(*user_service.UserModel)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	assertedComment.AuthorUserName = user.UserName
	// Check if the thread exist
	threadExistErr := orm.CheckCollectionForDocId(commentPayload.GetThreadId(), db.Collection("thread"), context.Background())
	if threadExistErr != nil {
		return nil, utilities.ReturnValidationError(threadExistErr)
	} 
	// Comment creation
	collection := orm.OrmSession.Client.Database("upvote").Collection("comment")
	commentId, err := orm.Create(assertedComment, collection, ctx)
	if err != nil {
		return nil, utilities.ReturnInternalError(err)
	}
	return &CreateCommentResponse{CommentId: commentId}, nil
}

// Streaming method to return the comments related to a certain thread
func (s *ThreadServer) GetThreadComments(req *GetThreadRequest, stream ThreadService_GetThreadCommentsServer) error{
	threadId := req.GetThreadId()
	oid, err := primitive.ObjectIDFromHex(threadId)
	if err != nil {
		return utilities.ReturnInternalError(err)
	}
	// Searching comments by thread
	collection := orm.OrmSession.Client.Database("upvote").Collection("comment")
	query := bson.M{"threadId": oid }
	cursor, err := collection.Find(context.Background(), query)
	if err != nil {
		return utilities.ReturnInternalError(err)
	}
	defer cursor.Close(context.Background())
	// Looping through the cursor
	for cursor.Next(context.Background()) {
		comment := CommentModel{}
		err := cursor.Decode(&comment)
		if err != nil {
			return utilities.ReturnInternalError(err)
		}
		commentResponse, err := orm.ConvertStruct(comment, Comment{})
		if err != nil {
			return utilities.ReturnInternalError(err)
		}
		stream.Send(&GetThreadCommentsResponse{
			Comment: commentResponse.(*Comment),
		})
	}
	return nil
}

