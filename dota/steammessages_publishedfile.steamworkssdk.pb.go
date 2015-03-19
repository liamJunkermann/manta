// Code generated by protoc-gen-go.
// source: steammessages_publishedfile.steamworkssdk.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import math "math"

// discarding unused import google_protobuf "github.com/dotabuff/manta/dota/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type CPublishedFile_Subscribe_Request struct {
	Publishedfileid  *uint64 `protobuf:"varint,1,opt,name=publishedfileid" json:"publishedfileid,omitempty"`
	Steamid          *uint64 `protobuf:"varint,2,opt,name=steamid" json:"steamid,omitempty"`
	ListType         *uint32 `protobuf:"varint,3,opt,name=list_type" json:"list_type,omitempty"`
	Appid            *int32  `protobuf:"varint,4,opt,name=appid" json:"appid,omitempty"`
	NotifyClient     *bool   `protobuf:"varint,5,opt,name=notify_client" json:"notify_client,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPublishedFile_Subscribe_Request) Reset()         { *m = CPublishedFile_Subscribe_Request{} }
func (m *CPublishedFile_Subscribe_Request) String() string { return proto.CompactTextString(m) }
func (*CPublishedFile_Subscribe_Request) ProtoMessage()    {}

func (m *CPublishedFile_Subscribe_Request) GetPublishedfileid() uint64 {
	if m != nil && m.Publishedfileid != nil {
		return *m.Publishedfileid
	}
	return 0
}

func (m *CPublishedFile_Subscribe_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CPublishedFile_Subscribe_Request) GetListType() uint32 {
	if m != nil && m.ListType != nil {
		return *m.ListType
	}
	return 0
}

func (m *CPublishedFile_Subscribe_Request) GetAppid() int32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CPublishedFile_Subscribe_Request) GetNotifyClient() bool {
	if m != nil && m.NotifyClient != nil {
		return *m.NotifyClient
	}
	return false
}

type CPublishedFile_Subscribe_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CPublishedFile_Subscribe_Response) Reset()         { *m = CPublishedFile_Subscribe_Response{} }
func (m *CPublishedFile_Subscribe_Response) String() string { return proto.CompactTextString(m) }
func (*CPublishedFile_Subscribe_Response) ProtoMessage()    {}

type CPublishedFile_Unsubscribe_Request struct {
	Publishedfileid  *uint64 `protobuf:"varint,1,opt,name=publishedfileid" json:"publishedfileid,omitempty"`
	Steamid          *uint64 `protobuf:"varint,2,opt,name=steamid" json:"steamid,omitempty"`
	ListType         *uint32 `protobuf:"varint,3,opt,name=list_type" json:"list_type,omitempty"`
	Appid            *int32  `protobuf:"varint,4,opt,name=appid" json:"appid,omitempty"`
	NotifyClient     *bool   `protobuf:"varint,5,opt,name=notify_client" json:"notify_client,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPublishedFile_Unsubscribe_Request) Reset()         { *m = CPublishedFile_Unsubscribe_Request{} }
func (m *CPublishedFile_Unsubscribe_Request) String() string { return proto.CompactTextString(m) }
func (*CPublishedFile_Unsubscribe_Request) ProtoMessage()    {}

func (m *CPublishedFile_Unsubscribe_Request) GetPublishedfileid() uint64 {
	if m != nil && m.Publishedfileid != nil {
		return *m.Publishedfileid
	}
	return 0
}

func (m *CPublishedFile_Unsubscribe_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CPublishedFile_Unsubscribe_Request) GetListType() uint32 {
	if m != nil && m.ListType != nil {
		return *m.ListType
	}
	return 0
}

func (m *CPublishedFile_Unsubscribe_Request) GetAppid() int32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CPublishedFile_Unsubscribe_Request) GetNotifyClient() bool {
	if m != nil && m.NotifyClient != nil {
		return *m.NotifyClient
	}
	return false
}

type CPublishedFile_Unsubscribe_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CPublishedFile_Unsubscribe_Response) Reset()         { *m = CPublishedFile_Unsubscribe_Response{} }
func (m *CPublishedFile_Unsubscribe_Response) String() string { return proto.CompactTextString(m) }
func (*CPublishedFile_Unsubscribe_Response) ProtoMessage()    {}

type CPublishedFile_Publish_Request struct {
	Appid                *uint32  `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	ConsumerAppid        *uint32  `protobuf:"varint,2,opt,name=consumer_appid" json:"consumer_appid,omitempty"`
	Cloudfilename        *string  `protobuf:"bytes,3,opt,name=cloudfilename" json:"cloudfilename,omitempty"`
	PreviewCloudfilename *string  `protobuf:"bytes,4,opt,name=preview_cloudfilename" json:"preview_cloudfilename,omitempty"`
	Title                *string  `protobuf:"bytes,5,opt,name=title" json:"title,omitempty"`
	FileDescription      *string  `protobuf:"bytes,6,opt,name=file_description" json:"file_description,omitempty"`
	FileType             *uint32  `protobuf:"varint,7,opt,name=file_type" json:"file_type,omitempty"`
	ConsumerShortcutName *string  `protobuf:"bytes,8,opt,name=consumer_shortcut_name" json:"consumer_shortcut_name,omitempty"`
	YoutubeUsername      *string  `protobuf:"bytes,9,opt,name=youtube_username" json:"youtube_username,omitempty"`
	YoutubeVideoid       *string  `protobuf:"bytes,10,opt,name=youtube_videoid" json:"youtube_videoid,omitempty"`
	Visibility           *uint32  `protobuf:"varint,11,opt,name=visibility" json:"visibility,omitempty"`
	RedirectUri          *string  `protobuf:"bytes,12,opt,name=redirect_uri" json:"redirect_uri,omitempty"`
	Tags                 []string `protobuf:"bytes,13,rep,name=tags" json:"tags,omitempty"`
	CollectionType       *string  `protobuf:"bytes,14,opt,name=collection_type" json:"collection_type,omitempty"`
	GameType             *string  `protobuf:"bytes,15,opt,name=game_type" json:"game_type,omitempty"`
	Url                  *string  `protobuf:"bytes,16,opt,name=url" json:"url,omitempty"`
	XXX_unrecognized     []byte   `json:"-"`
}

func (m *CPublishedFile_Publish_Request) Reset()         { *m = CPublishedFile_Publish_Request{} }
func (m *CPublishedFile_Publish_Request) String() string { return proto.CompactTextString(m) }
func (*CPublishedFile_Publish_Request) ProtoMessage()    {}

func (m *CPublishedFile_Publish_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CPublishedFile_Publish_Request) GetConsumerAppid() uint32 {
	if m != nil && m.ConsumerAppid != nil {
		return *m.ConsumerAppid
	}
	return 0
}

func (m *CPublishedFile_Publish_Request) GetCloudfilename() string {
	if m != nil && m.Cloudfilename != nil {
		return *m.Cloudfilename
	}
	return ""
}

func (m *CPublishedFile_Publish_Request) GetPreviewCloudfilename() string {
	if m != nil && m.PreviewCloudfilename != nil {
		return *m.PreviewCloudfilename
	}
	return ""
}

func (m *CPublishedFile_Publish_Request) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *CPublishedFile_Publish_Request) GetFileDescription() string {
	if m != nil && m.FileDescription != nil {
		return *m.FileDescription
	}
	return ""
}

func (m *CPublishedFile_Publish_Request) GetFileType() uint32 {
	if m != nil && m.FileType != nil {
		return *m.FileType
	}
	return 0
}

func (m *CPublishedFile_Publish_Request) GetConsumerShortcutName() string {
	if m != nil && m.ConsumerShortcutName != nil {
		return *m.ConsumerShortcutName
	}
	return ""
}

func (m *CPublishedFile_Publish_Request) GetYoutubeUsername() string {
	if m != nil && m.YoutubeUsername != nil {
		return *m.YoutubeUsername
	}
	return ""
}

func (m *CPublishedFile_Publish_Request) GetYoutubeVideoid() string {
	if m != nil && m.YoutubeVideoid != nil {
		return *m.YoutubeVideoid
	}
	return ""
}

func (m *CPublishedFile_Publish_Request) GetVisibility() uint32 {
	if m != nil && m.Visibility != nil {
		return *m.Visibility
	}
	return 0
}

func (m *CPublishedFile_Publish_Request) GetRedirectUri() string {
	if m != nil && m.RedirectUri != nil {
		return *m.RedirectUri
	}
	return ""
}

func (m *CPublishedFile_Publish_Request) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *CPublishedFile_Publish_Request) GetCollectionType() string {
	if m != nil && m.CollectionType != nil {
		return *m.CollectionType
	}
	return ""
}

func (m *CPublishedFile_Publish_Request) GetGameType() string {
	if m != nil && m.GameType != nil {
		return *m.GameType
	}
	return ""
}

func (m *CPublishedFile_Publish_Request) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

type CPublishedFile_Publish_Response struct {
	Publishedfileid  *uint64 `protobuf:"varint,1,opt,name=publishedfileid" json:"publishedfileid,omitempty"`
	RedirectUri      *string `protobuf:"bytes,2,opt,name=redirect_uri" json:"redirect_uri,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPublishedFile_Publish_Response) Reset()         { *m = CPublishedFile_Publish_Response{} }
func (m *CPublishedFile_Publish_Response) String() string { return proto.CompactTextString(m) }
func (*CPublishedFile_Publish_Response) ProtoMessage()    {}

func (m *CPublishedFile_Publish_Response) GetPublishedfileid() uint64 {
	if m != nil && m.Publishedfileid != nil {
		return *m.Publishedfileid
	}
	return 0
}

func (m *CPublishedFile_Publish_Response) GetRedirectUri() string {
	if m != nil && m.RedirectUri != nil {
		return *m.RedirectUri
	}
	return ""
}

type CPublishedFile_GetDetails_Request struct {
	Publishedfileids          []uint64 `protobuf:"fixed64,1,rep,name=publishedfileids" json:"publishedfileids,omitempty"`
	Includetags               *bool    `protobuf:"varint,2,opt,name=includetags" json:"includetags,omitempty"`
	Includeadditionalpreviews *bool    `protobuf:"varint,3,opt,name=includeadditionalpreviews" json:"includeadditionalpreviews,omitempty"`
	Includechildren           *bool    `protobuf:"varint,4,opt,name=includechildren" json:"includechildren,omitempty"`
	Appid                     *uint32  `protobuf:"varint,5,opt,name=appid" json:"appid,omitempty"`
	XXX_unrecognized          []byte   `json:"-"`
}

func (m *CPublishedFile_GetDetails_Request) Reset()         { *m = CPublishedFile_GetDetails_Request{} }
func (m *CPublishedFile_GetDetails_Request) String() string { return proto.CompactTextString(m) }
func (*CPublishedFile_GetDetails_Request) ProtoMessage()    {}

func (m *CPublishedFile_GetDetails_Request) GetPublishedfileids() []uint64 {
	if m != nil {
		return m.Publishedfileids
	}
	return nil
}

func (m *CPublishedFile_GetDetails_Request) GetIncludetags() bool {
	if m != nil && m.Includetags != nil {
		return *m.Includetags
	}
	return false
}

func (m *CPublishedFile_GetDetails_Request) GetIncludeadditionalpreviews() bool {
	if m != nil && m.Includeadditionalpreviews != nil {
		return *m.Includeadditionalpreviews
	}
	return false
}

func (m *CPublishedFile_GetDetails_Request) GetIncludechildren() bool {
	if m != nil && m.Includechildren != nil {
		return *m.Includechildren
	}
	return false
}

func (m *CPublishedFile_GetDetails_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

type PublishedFileDetails struct {
	Result                      *uint32                         `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Publishedfileid             *uint64                         `protobuf:"varint,2,opt,name=publishedfileid" json:"publishedfileid,omitempty"`
	Creator                     *uint64                         `protobuf:"fixed64,3,opt,name=creator" json:"creator,omitempty"`
	CreatorAppid                *uint32                         `protobuf:"varint,4,opt,name=creator_appid" json:"creator_appid,omitempty"`
	ConsumerAppid               *uint32                         `protobuf:"varint,5,opt,name=consumer_appid" json:"consumer_appid,omitempty"`
	ConsumerShortcutid          *uint32                         `protobuf:"varint,6,opt,name=consumer_shortcutid" json:"consumer_shortcutid,omitempty"`
	Filename                    *string                         `protobuf:"bytes,7,opt,name=filename" json:"filename,omitempty"`
	FileSize                    *uint64                         `protobuf:"varint,8,opt,name=file_size" json:"file_size,omitempty"`
	FileUrl                     *string                         `protobuf:"bytes,9,opt,name=file_url" json:"file_url,omitempty"`
	PreviewUrl                  *string                         `protobuf:"bytes,10,opt,name=preview_url" json:"preview_url,omitempty"`
	Youtubevideoid              *string                         `protobuf:"bytes,11,opt,name=youtubevideoid" json:"youtubevideoid,omitempty"`
	Url                         *string                         `protobuf:"bytes,12,opt,name=url" json:"url,omitempty"`
	HcontentFile                *uint64                         `protobuf:"fixed64,13,opt,name=hcontent_file" json:"hcontent_file,omitempty"`
	HcontentPreview             *uint64                         `protobuf:"fixed64,14,opt,name=hcontent_preview" json:"hcontent_preview,omitempty"`
	Title                       *string                         `protobuf:"bytes,15,opt,name=title" json:"title,omitempty"`
	FileDescription             *string                         `protobuf:"bytes,16,opt,name=file_description" json:"file_description,omitempty"`
	TimeCreated                 *uint32                         `protobuf:"varint,17,opt,name=time_created" json:"time_created,omitempty"`
	TimeUpdated                 *uint32                         `protobuf:"varint,18,opt,name=time_updated" json:"time_updated,omitempty"`
	Visibility                  *uint32                         `protobuf:"varint,19,opt,name=visibility" json:"visibility,omitempty"`
	IsFriendOfOwner             *bool                           `protobuf:"varint,20,opt,name=is_friend_of_owner" json:"is_friend_of_owner,omitempty"`
	Flags                       *uint32                         `protobuf:"varint,21,opt,name=flags" json:"flags,omitempty"`
	WorkshopFile                *bool                           `protobuf:"varint,22,opt,name=workshop_file" json:"workshop_file,omitempty"`
	WorkshopAccepted            *bool                           `protobuf:"varint,23,opt,name=workshop_accepted" json:"workshop_accepted,omitempty"`
	ShowSubscribeAll            *bool                           `protobuf:"varint,24,opt,name=show_subscribe_all" json:"show_subscribe_all,omitempty"`
	NumCommentsDeveloper        *int32                          `protobuf:"varint,25,opt,name=num_comments_developer" json:"num_comments_developer,omitempty"`
	NumCommentsPublic           *int32                          `protobuf:"varint,26,opt,name=num_comments_public" json:"num_comments_public,omitempty"`
	Banned                      *bool                           `protobuf:"varint,27,opt,name=banned" json:"banned,omitempty"`
	BanReason                   *string                         `protobuf:"bytes,28,opt,name=ban_reason" json:"ban_reason,omitempty"`
	Banner                      *uint64                         `protobuf:"fixed64,29,opt,name=banner" json:"banner,omitempty"`
	CanBeDeleted                *bool                           `protobuf:"varint,30,opt,name=can_be_deleted" json:"can_be_deleted,omitempty"`
	ConsumerAppPrivate          *bool                           `protobuf:"varint,31,opt,name=consumer_app_private" json:"consumer_app_private,omitempty"`
	Incompatible                *bool                           `protobuf:"varint,32,opt,name=incompatible" json:"incompatible,omitempty"`
	ConsumerAppInstructionsLink *string                         `protobuf:"bytes,33,opt,name=consumer_app_instructions_link" json:"consumer_app_instructions_link,omitempty"`
	ConsumerAppName             *string                         `protobuf:"bytes,34,opt,name=consumer_app_name" json:"consumer_app_name,omitempty"`
	FileType                    *uint32                         `protobuf:"varint,35,opt,name=file_type" json:"file_type,omitempty"`
	CanSubscribe                *bool                           `protobuf:"varint,36,opt,name=can_subscribe" json:"can_subscribe,omitempty"`
	Subscriptions               *uint32                         `protobuf:"varint,37,opt,name=subscriptions" json:"subscriptions,omitempty"`
	Favorited                   *uint32                         `protobuf:"varint,38,opt,name=favorited" json:"favorited,omitempty"`
	LifetimeSubscriptions       *uint32                         `protobuf:"varint,39,opt,name=lifetime_subscriptions" json:"lifetime_subscriptions,omitempty"`
	LifetimeFavorited           *uint32                         `protobuf:"varint,40,opt,name=lifetime_favorited" json:"lifetime_favorited,omitempty"`
	Views                       *uint32                         `protobuf:"varint,41,opt,name=views" json:"views,omitempty"`
	ImageWidth                  *uint32                         `protobuf:"varint,42,opt,name=image_width" json:"image_width,omitempty"`
	ImageHeight                 *uint32                         `protobuf:"varint,43,opt,name=image_height" json:"image_height,omitempty"`
	SpoilerTag                  *bool                           `protobuf:"varint,44,opt,name=spoiler_tag" json:"spoiler_tag,omitempty"`
	Shortcutid                  *uint32                         `protobuf:"varint,45,opt,name=shortcutid" json:"shortcutid,omitempty"`
	Shortcutname                *string                         `protobuf:"bytes,46,opt,name=shortcutname" json:"shortcutname,omitempty"`
	ShortDescription            *string                         `protobuf:"bytes,47,opt,name=short_description" json:"short_description,omitempty"`
	ImageUrl                    *string                         `protobuf:"bytes,48,opt,name=image_url" json:"image_url,omitempty"`
	AppName                     *string                         `protobuf:"bytes,49,opt,name=app_name" json:"app_name,omitempty"`
	TimeSubscribed              *uint32                         `protobuf:"varint,50,opt,name=time_subscribed" json:"time_subscribed,omitempty"`
	Previews                    []*PublishedFileDetails_Preview `protobuf:"bytes,51,rep,name=previews" json:"previews,omitempty"`
	Tags                        []*PublishedFileDetails_Tag     `protobuf:"bytes,52,rep,name=tags" json:"tags,omitempty"`
	Children                    []*PublishedFileDetails_Child   `protobuf:"bytes,53,rep,name=children" json:"children,omitempty"`
	Kvtags                      []*PublishedFileDetails_KVTag   `protobuf:"bytes,54,rep,name=kvtags" json:"kvtags,omitempty"`
	XXX_unrecognized            []byte                          `json:"-"`
}

func (m *PublishedFileDetails) Reset()         { *m = PublishedFileDetails{} }
func (m *PublishedFileDetails) String() string { return proto.CompactTextString(m) }
func (*PublishedFileDetails) ProtoMessage()    {}

func (m *PublishedFileDetails) GetResult() uint32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *PublishedFileDetails) GetPublishedfileid() uint64 {
	if m != nil && m.Publishedfileid != nil {
		return *m.Publishedfileid
	}
	return 0
}

func (m *PublishedFileDetails) GetCreator() uint64 {
	if m != nil && m.Creator != nil {
		return *m.Creator
	}
	return 0
}

func (m *PublishedFileDetails) GetCreatorAppid() uint32 {
	if m != nil && m.CreatorAppid != nil {
		return *m.CreatorAppid
	}
	return 0
}

func (m *PublishedFileDetails) GetConsumerAppid() uint32 {
	if m != nil && m.ConsumerAppid != nil {
		return *m.ConsumerAppid
	}
	return 0
}

func (m *PublishedFileDetails) GetConsumerShortcutid() uint32 {
	if m != nil && m.ConsumerShortcutid != nil {
		return *m.ConsumerShortcutid
	}
	return 0
}

func (m *PublishedFileDetails) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *PublishedFileDetails) GetFileSize() uint64 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *PublishedFileDetails) GetFileUrl() string {
	if m != nil && m.FileUrl != nil {
		return *m.FileUrl
	}
	return ""
}

func (m *PublishedFileDetails) GetPreviewUrl() string {
	if m != nil && m.PreviewUrl != nil {
		return *m.PreviewUrl
	}
	return ""
}

func (m *PublishedFileDetails) GetYoutubevideoid() string {
	if m != nil && m.Youtubevideoid != nil {
		return *m.Youtubevideoid
	}
	return ""
}

func (m *PublishedFileDetails) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *PublishedFileDetails) GetHcontentFile() uint64 {
	if m != nil && m.HcontentFile != nil {
		return *m.HcontentFile
	}
	return 0
}

func (m *PublishedFileDetails) GetHcontentPreview() uint64 {
	if m != nil && m.HcontentPreview != nil {
		return *m.HcontentPreview
	}
	return 0
}

func (m *PublishedFileDetails) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *PublishedFileDetails) GetFileDescription() string {
	if m != nil && m.FileDescription != nil {
		return *m.FileDescription
	}
	return ""
}

func (m *PublishedFileDetails) GetTimeCreated() uint32 {
	if m != nil && m.TimeCreated != nil {
		return *m.TimeCreated
	}
	return 0
}

func (m *PublishedFileDetails) GetTimeUpdated() uint32 {
	if m != nil && m.TimeUpdated != nil {
		return *m.TimeUpdated
	}
	return 0
}

func (m *PublishedFileDetails) GetVisibility() uint32 {
	if m != nil && m.Visibility != nil {
		return *m.Visibility
	}
	return 0
}

func (m *PublishedFileDetails) GetIsFriendOfOwner() bool {
	if m != nil && m.IsFriendOfOwner != nil {
		return *m.IsFriendOfOwner
	}
	return false
}

func (m *PublishedFileDetails) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *PublishedFileDetails) GetWorkshopFile() bool {
	if m != nil && m.WorkshopFile != nil {
		return *m.WorkshopFile
	}
	return false
}

func (m *PublishedFileDetails) GetWorkshopAccepted() bool {
	if m != nil && m.WorkshopAccepted != nil {
		return *m.WorkshopAccepted
	}
	return false
}

func (m *PublishedFileDetails) GetShowSubscribeAll() bool {
	if m != nil && m.ShowSubscribeAll != nil {
		return *m.ShowSubscribeAll
	}
	return false
}

func (m *PublishedFileDetails) GetNumCommentsDeveloper() int32 {
	if m != nil && m.NumCommentsDeveloper != nil {
		return *m.NumCommentsDeveloper
	}
	return 0
}

func (m *PublishedFileDetails) GetNumCommentsPublic() int32 {
	if m != nil && m.NumCommentsPublic != nil {
		return *m.NumCommentsPublic
	}
	return 0
}

func (m *PublishedFileDetails) GetBanned() bool {
	if m != nil && m.Banned != nil {
		return *m.Banned
	}
	return false
}

func (m *PublishedFileDetails) GetBanReason() string {
	if m != nil && m.BanReason != nil {
		return *m.BanReason
	}
	return ""
}

func (m *PublishedFileDetails) GetBanner() uint64 {
	if m != nil && m.Banner != nil {
		return *m.Banner
	}
	return 0
}

func (m *PublishedFileDetails) GetCanBeDeleted() bool {
	if m != nil && m.CanBeDeleted != nil {
		return *m.CanBeDeleted
	}
	return false
}

func (m *PublishedFileDetails) GetConsumerAppPrivate() bool {
	if m != nil && m.ConsumerAppPrivate != nil {
		return *m.ConsumerAppPrivate
	}
	return false
}

func (m *PublishedFileDetails) GetIncompatible() bool {
	if m != nil && m.Incompatible != nil {
		return *m.Incompatible
	}
	return false
}

func (m *PublishedFileDetails) GetConsumerAppInstructionsLink() string {
	if m != nil && m.ConsumerAppInstructionsLink != nil {
		return *m.ConsumerAppInstructionsLink
	}
	return ""
}

func (m *PublishedFileDetails) GetConsumerAppName() string {
	if m != nil && m.ConsumerAppName != nil {
		return *m.ConsumerAppName
	}
	return ""
}

func (m *PublishedFileDetails) GetFileType() uint32 {
	if m != nil && m.FileType != nil {
		return *m.FileType
	}
	return 0
}

func (m *PublishedFileDetails) GetCanSubscribe() bool {
	if m != nil && m.CanSubscribe != nil {
		return *m.CanSubscribe
	}
	return false
}

func (m *PublishedFileDetails) GetSubscriptions() uint32 {
	if m != nil && m.Subscriptions != nil {
		return *m.Subscriptions
	}
	return 0
}

func (m *PublishedFileDetails) GetFavorited() uint32 {
	if m != nil && m.Favorited != nil {
		return *m.Favorited
	}
	return 0
}

func (m *PublishedFileDetails) GetLifetimeSubscriptions() uint32 {
	if m != nil && m.LifetimeSubscriptions != nil {
		return *m.LifetimeSubscriptions
	}
	return 0
}

func (m *PublishedFileDetails) GetLifetimeFavorited() uint32 {
	if m != nil && m.LifetimeFavorited != nil {
		return *m.LifetimeFavorited
	}
	return 0
}

func (m *PublishedFileDetails) GetViews() uint32 {
	if m != nil && m.Views != nil {
		return *m.Views
	}
	return 0
}

func (m *PublishedFileDetails) GetImageWidth() uint32 {
	if m != nil && m.ImageWidth != nil {
		return *m.ImageWidth
	}
	return 0
}

func (m *PublishedFileDetails) GetImageHeight() uint32 {
	if m != nil && m.ImageHeight != nil {
		return *m.ImageHeight
	}
	return 0
}

func (m *PublishedFileDetails) GetSpoilerTag() bool {
	if m != nil && m.SpoilerTag != nil {
		return *m.SpoilerTag
	}
	return false
}

func (m *PublishedFileDetails) GetShortcutid() uint32 {
	if m != nil && m.Shortcutid != nil {
		return *m.Shortcutid
	}
	return 0
}

func (m *PublishedFileDetails) GetShortcutname() string {
	if m != nil && m.Shortcutname != nil {
		return *m.Shortcutname
	}
	return ""
}

func (m *PublishedFileDetails) GetShortDescription() string {
	if m != nil && m.ShortDescription != nil {
		return *m.ShortDescription
	}
	return ""
}

func (m *PublishedFileDetails) GetImageUrl() string {
	if m != nil && m.ImageUrl != nil {
		return *m.ImageUrl
	}
	return ""
}

func (m *PublishedFileDetails) GetAppName() string {
	if m != nil && m.AppName != nil {
		return *m.AppName
	}
	return ""
}

func (m *PublishedFileDetails) GetTimeSubscribed() uint32 {
	if m != nil && m.TimeSubscribed != nil {
		return *m.TimeSubscribed
	}
	return 0
}

func (m *PublishedFileDetails) GetPreviews() []*PublishedFileDetails_Preview {
	if m != nil {
		return m.Previews
	}
	return nil
}

func (m *PublishedFileDetails) GetTags() []*PublishedFileDetails_Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *PublishedFileDetails) GetChildren() []*PublishedFileDetails_Child {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *PublishedFileDetails) GetKvtags() []*PublishedFileDetails_KVTag {
	if m != nil {
		return m.Kvtags
	}
	return nil
}

type PublishedFileDetails_Tag struct {
	Tag              *string `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
	Adminonly        *bool   `protobuf:"varint,2,opt,name=adminonly" json:"adminonly,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PublishedFileDetails_Tag) Reset()         { *m = PublishedFileDetails_Tag{} }
func (m *PublishedFileDetails_Tag) String() string { return proto.CompactTextString(m) }
func (*PublishedFileDetails_Tag) ProtoMessage()    {}

func (m *PublishedFileDetails_Tag) GetTag() string {
	if m != nil && m.Tag != nil {
		return *m.Tag
	}
	return ""
}

func (m *PublishedFileDetails_Tag) GetAdminonly() bool {
	if m != nil && m.Adminonly != nil {
		return *m.Adminonly
	}
	return false
}

type PublishedFileDetails_Preview struct {
	Previewid        *uint64 `protobuf:"varint,1,opt,name=previewid" json:"previewid,omitempty"`
	Sortorder        *uint32 `protobuf:"varint,2,opt,name=sortorder" json:"sortorder,omitempty"`
	Url              *string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Size             *uint32 `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
	Filename         *string `protobuf:"bytes,5,opt,name=filename" json:"filename,omitempty"`
	Youtubevideoid   *string `protobuf:"bytes,6,opt,name=youtubevideoid" json:"youtubevideoid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PublishedFileDetails_Preview) Reset()         { *m = PublishedFileDetails_Preview{} }
func (m *PublishedFileDetails_Preview) String() string { return proto.CompactTextString(m) }
func (*PublishedFileDetails_Preview) ProtoMessage()    {}

func (m *PublishedFileDetails_Preview) GetPreviewid() uint64 {
	if m != nil && m.Previewid != nil {
		return *m.Previewid
	}
	return 0
}

func (m *PublishedFileDetails_Preview) GetSortorder() uint32 {
	if m != nil && m.Sortorder != nil {
		return *m.Sortorder
	}
	return 0
}

func (m *PublishedFileDetails_Preview) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *PublishedFileDetails_Preview) GetSize() uint32 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *PublishedFileDetails_Preview) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *PublishedFileDetails_Preview) GetYoutubevideoid() string {
	if m != nil && m.Youtubevideoid != nil {
		return *m.Youtubevideoid
	}
	return ""
}

type PublishedFileDetails_Child struct {
	Publishedfileid  *uint64 `protobuf:"varint,1,opt,name=publishedfileid" json:"publishedfileid,omitempty"`
	Sortorder        *uint32 `protobuf:"varint,2,opt,name=sortorder" json:"sortorder,omitempty"`
	FileType         *uint32 `protobuf:"varint,3,opt,name=file_type" json:"file_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PublishedFileDetails_Child) Reset()         { *m = PublishedFileDetails_Child{} }
func (m *PublishedFileDetails_Child) String() string { return proto.CompactTextString(m) }
func (*PublishedFileDetails_Child) ProtoMessage()    {}

func (m *PublishedFileDetails_Child) GetPublishedfileid() uint64 {
	if m != nil && m.Publishedfileid != nil {
		return *m.Publishedfileid
	}
	return 0
}

func (m *PublishedFileDetails_Child) GetSortorder() uint32 {
	if m != nil && m.Sortorder != nil {
		return *m.Sortorder
	}
	return 0
}

func (m *PublishedFileDetails_Child) GetFileType() uint32 {
	if m != nil && m.FileType != nil {
		return *m.FileType
	}
	return 0
}

type PublishedFileDetails_KVTag struct {
	Key              *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PublishedFileDetails_KVTag) Reset()         { *m = PublishedFileDetails_KVTag{} }
func (m *PublishedFileDetails_KVTag) String() string { return proto.CompactTextString(m) }
func (*PublishedFileDetails_KVTag) ProtoMessage()    {}

func (m *PublishedFileDetails_KVTag) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *PublishedFileDetails_KVTag) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type CPublishedFile_GetDetails_Response struct {
	Publishedfiledetails []*PublishedFileDetails `protobuf:"bytes,1,rep,name=publishedfiledetails" json:"publishedfiledetails,omitempty"`
	XXX_unrecognized     []byte                  `json:"-"`
}

func (m *CPublishedFile_GetDetails_Response) Reset()         { *m = CPublishedFile_GetDetails_Response{} }
func (m *CPublishedFile_GetDetails_Response) String() string { return proto.CompactTextString(m) }
func (*CPublishedFile_GetDetails_Response) ProtoMessage()    {}

func (m *CPublishedFile_GetDetails_Response) GetPublishedfiledetails() []*PublishedFileDetails {
	if m != nil {
		return m.Publishedfiledetails
	}
	return nil
}

type CPublishedFile_GetUserFiles_Request struct {
	Appid            *uint32  `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Page             *uint32  `protobuf:"varint,3,opt,name=page,def=1" json:"page,omitempty"`
	Numperpage       *uint32  `protobuf:"varint,4,opt,name=numperpage,def=1" json:"numperpage,omitempty"`
	Sortmethod       *string  `protobuf:"bytes,6,opt,name=sortmethod,def=lastupdated" json:"sortmethod,omitempty"`
	Totalonly        *bool    `protobuf:"varint,7,opt,name=totalonly" json:"totalonly,omitempty"`
	Privacy          *uint32  `protobuf:"varint,9,opt,name=privacy" json:"privacy,omitempty"`
	IdsOnly          *bool    `protobuf:"varint,10,opt,name=ids_only" json:"ids_only,omitempty"`
	Requiredtags     []string `protobuf:"bytes,11,rep,name=requiredtags" json:"requiredtags,omitempty"`
	Excludedtags     []string `protobuf:"bytes,12,rep,name=excludedtags" json:"excludedtags,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CPublishedFile_GetUserFiles_Request) Reset()         { *m = CPublishedFile_GetUserFiles_Request{} }
func (m *CPublishedFile_GetUserFiles_Request) String() string { return proto.CompactTextString(m) }
func (*CPublishedFile_GetUserFiles_Request) ProtoMessage()    {}

const Default_CPublishedFile_GetUserFiles_Request_Page uint32 = 1
const Default_CPublishedFile_GetUserFiles_Request_Numperpage uint32 = 1
const Default_CPublishedFile_GetUserFiles_Request_Sortmethod string = "lastupdated"

func (m *CPublishedFile_GetUserFiles_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CPublishedFile_GetUserFiles_Request) GetPage() uint32 {
	if m != nil && m.Page != nil {
		return *m.Page
	}
	return Default_CPublishedFile_GetUserFiles_Request_Page
}

func (m *CPublishedFile_GetUserFiles_Request) GetNumperpage() uint32 {
	if m != nil && m.Numperpage != nil {
		return *m.Numperpage
	}
	return Default_CPublishedFile_GetUserFiles_Request_Numperpage
}

func (m *CPublishedFile_GetUserFiles_Request) GetSortmethod() string {
	if m != nil && m.Sortmethod != nil {
		return *m.Sortmethod
	}
	return Default_CPublishedFile_GetUserFiles_Request_Sortmethod
}

func (m *CPublishedFile_GetUserFiles_Request) GetTotalonly() bool {
	if m != nil && m.Totalonly != nil {
		return *m.Totalonly
	}
	return false
}

func (m *CPublishedFile_GetUserFiles_Request) GetPrivacy() uint32 {
	if m != nil && m.Privacy != nil {
		return *m.Privacy
	}
	return 0
}

func (m *CPublishedFile_GetUserFiles_Request) GetIdsOnly() bool {
	if m != nil && m.IdsOnly != nil {
		return *m.IdsOnly
	}
	return false
}

func (m *CPublishedFile_GetUserFiles_Request) GetRequiredtags() []string {
	if m != nil {
		return m.Requiredtags
	}
	return nil
}

func (m *CPublishedFile_GetUserFiles_Request) GetExcludedtags() []string {
	if m != nil {
		return m.Excludedtags
	}
	return nil
}

type CPublishedFile_GetUserFiles_Response struct {
	Total                *uint32                                     `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
	Startindex           *uint32                                     `protobuf:"varint,2,opt,name=startindex" json:"startindex,omitempty"`
	Publishedfiledetails []*PublishedFileDetails                     `protobuf:"bytes,3,rep,name=publishedfiledetails" json:"publishedfiledetails,omitempty"`
	Apps                 []*CPublishedFile_GetUserFiles_Response_App `protobuf:"bytes,4,rep,name=apps" json:"apps,omitempty"`
	XXX_unrecognized     []byte                                      `json:"-"`
}

func (m *CPublishedFile_GetUserFiles_Response) Reset()         { *m = CPublishedFile_GetUserFiles_Response{} }
func (m *CPublishedFile_GetUserFiles_Response) String() string { return proto.CompactTextString(m) }
func (*CPublishedFile_GetUserFiles_Response) ProtoMessage()    {}

func (m *CPublishedFile_GetUserFiles_Response) GetTotal() uint32 {
	if m != nil && m.Total != nil {
		return *m.Total
	}
	return 0
}

func (m *CPublishedFile_GetUserFiles_Response) GetStartindex() uint32 {
	if m != nil && m.Startindex != nil {
		return *m.Startindex
	}
	return 0
}

func (m *CPublishedFile_GetUserFiles_Response) GetPublishedfiledetails() []*PublishedFileDetails {
	if m != nil {
		return m.Publishedfiledetails
	}
	return nil
}

func (m *CPublishedFile_GetUserFiles_Response) GetApps() []*CPublishedFile_GetUserFiles_Response_App {
	if m != nil {
		return m.Apps
	}
	return nil
}

type CPublishedFile_GetUserFiles_Response_App struct {
	Appid            *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Shortcutid       *uint32 `protobuf:"varint,3,opt,name=shortcutid" json:"shortcutid,omitempty"`
	Private          *bool   `protobuf:"varint,4,opt,name=private" json:"private,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPublishedFile_GetUserFiles_Response_App) Reset() {
	*m = CPublishedFile_GetUserFiles_Response_App{}
}
func (m *CPublishedFile_GetUserFiles_Response_App) String() string { return proto.CompactTextString(m) }
func (*CPublishedFile_GetUserFiles_Response_App) ProtoMessage()    {}

func (m *CPublishedFile_GetUserFiles_Response_App) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CPublishedFile_GetUserFiles_Response_App) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CPublishedFile_GetUserFiles_Response_App) GetShortcutid() uint32 {
	if m != nil && m.Shortcutid != nil {
		return *m.Shortcutid
	}
	return 0
}

func (m *CPublishedFile_GetUserFiles_Response_App) GetPrivate() bool {
	if m != nil && m.Private != nil {
		return *m.Private
	}
	return false
}

type CPublishedFile_Update_Request struct {
	Appid            *uint32  `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Publishedfileid  *uint64  `protobuf:"fixed64,2,opt,name=publishedfileid" json:"publishedfileid,omitempty"`
	Title            *string  `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	FileDescription  *string  `protobuf:"bytes,4,opt,name=file_description" json:"file_description,omitempty"`
	Visibility       *uint32  `protobuf:"varint,5,opt,name=visibility" json:"visibility,omitempty"`
	Tags             []string `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	Filename         *string  `protobuf:"bytes,7,opt,name=filename" json:"filename,omitempty"`
	PreviewFilename  *string  `protobuf:"bytes,8,opt,name=preview_filename" json:"preview_filename,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CPublishedFile_Update_Request) Reset()         { *m = CPublishedFile_Update_Request{} }
func (m *CPublishedFile_Update_Request) String() string { return proto.CompactTextString(m) }
func (*CPublishedFile_Update_Request) ProtoMessage()    {}

func (m *CPublishedFile_Update_Request) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CPublishedFile_Update_Request) GetPublishedfileid() uint64 {
	if m != nil && m.Publishedfileid != nil {
		return *m.Publishedfileid
	}
	return 0
}

func (m *CPublishedFile_Update_Request) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *CPublishedFile_Update_Request) GetFileDescription() string {
	if m != nil && m.FileDescription != nil {
		return *m.FileDescription
	}
	return ""
}

func (m *CPublishedFile_Update_Request) GetVisibility() uint32 {
	if m != nil && m.Visibility != nil {
		return *m.Visibility
	}
	return 0
}

func (m *CPublishedFile_Update_Request) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *CPublishedFile_Update_Request) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CPublishedFile_Update_Request) GetPreviewFilename() string {
	if m != nil && m.PreviewFilename != nil {
		return *m.PreviewFilename
	}
	return ""
}

type CPublishedFile_Update_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CPublishedFile_Update_Response) Reset()         { *m = CPublishedFile_Update_Response{} }
func (m *CPublishedFile_Update_Response) String() string { return proto.CompactTextString(m) }
func (*CPublishedFile_Update_Response) ProtoMessage()    {}

func init() {
}
