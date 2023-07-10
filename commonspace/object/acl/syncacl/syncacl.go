package syncacl

import (
	"context"

	"github.com/anyproto/any-sync/accountservice"
	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/commonspace/object/acl/list"
	"github.com/anyproto/any-sync/commonspace/spacestorage"
	"github.com/anyproto/any-sync/commonspace/spacesyncproto"
)

const CName = "common.acl.syncacl"

func New() *SyncAcl {
	return &SyncAcl{}
}

type SyncAcl struct {
	list.AclList
}

func (s *SyncAcl) HandleRequest(ctx context.Context, senderId string, request *spacesyncproto.ObjectSyncMessage) (response *spacesyncproto.ObjectSyncMessage, err error) {
	return nil, nil
}

func (s *SyncAcl) HandleMessage(ctx context.Context, senderId string, request *spacesyncproto.ObjectSyncMessage) (err error) {
	return nil
}

func (s *SyncAcl) Init(a *app.App) (err error) {
	storage := a.MustComponent(spacestorage.CName).(spacestorage.SpaceStorage)
	aclStorage, err := storage.AclStorage()
	if err != nil {
		return err
	}
	acc := a.MustComponent(accountservice.CName).(accountservice.Service)
	s.AclList, err = list.BuildAclListWithIdentity(acc.Account(), aclStorage, list.NoOpAcceptorVerifier{})
	return err
}

func (s *SyncAcl) Name() (name string) {
	return CName
}
