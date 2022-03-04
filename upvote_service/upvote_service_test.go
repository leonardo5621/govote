package upvote_service

import (
	"testing"
	"gotest.tools/v3/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

const bufSize = 1024 * 1024

//var lis *bufconn.Listener

func TestVotehandler(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("Vote handler test", func(mt *mtest.T) {
		collection := mt.Coll
		oid := primitive.NewObjectID()
		//Down vote again
		newVote := &UpvoteThreadModel{
			Id: &oid,
			Votedir: int32(-1),
		}
		existingVote := bson.M{
			"_id": oid,
			"votedir": int32(-1),
		}
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		downVoteUpdateFactor, err := manageExistingVote(newVote, collection, existingVote)
		if err != nil {
			t.Error(err)
		}
		var expectedDownvoteFactor int32 = 1
		assert.Equal(t, downVoteUpdateFactor, expectedDownvoteFactor)

		//Upvote again
		newVote.Votedir = int32(1)
		existingVote["votedir"] = int32(1)
		upVoteUpdateFactor, err := manageExistingVote(newVote, collection, existingVote)
		if err != nil {
			t.Error(err)
		}
		var expectedUpvoteFactor int32 = -1
		assert.Equal(t, upVoteUpdateFactor, expectedUpvoteFactor)

		//Revert downvote
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		newVote.Votedir = int32(1)
		existingVote["votedir"] = int32(-1)
		revertDownVoteUpdateFactor, err := manageExistingVote(newVote, collection, existingVote)
		if err != nil {
			t.Error(err)
		}
		var expectedRevertDownvoteFactor int32 = 2
		assert.Equal(t, revertDownVoteUpdateFactor, expectedRevertDownvoteFactor)

		//Revert upvote
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		newVote.Votedir = int32(-1)
		existingVote["votedir"] = int32(1)
		revertUpVoteUpdateFactor, err := manageExistingVote(newVote, collection, existingVote)
		if err != nil {
			t.Error(err)
		}
		var expectedRevertUpvoteFactor int32 = -2
		assert.Equal(t, revertUpVoteUpdateFactor, expectedRevertUpvoteFactor)
	})
}