package customers

import (
	"context"
	"github.com/integrica-io/iManCloudCore/internal"
	"github.com/integrica-io/iManCloudCore/client"	
)

type GetRoleDetailsOutput struct {
	Data struct {
		Description string `json:"description"`
		Database    string `json:"database"`
		ID          string `json:"id"`
		M1          int    `json:"m1"`
		M1Bits      struct {
			EditExternalDefaultSecurityFolder bool `json:"edit_external_default_security_folder"`
			DeletePublicFolder                bool `json:"delete_public_folder"`
			EditExternalDefaultSecurity       bool `json:"edit_external_default_security"`
			Checkout                          bool `json:"checkout"`
			AllowIndexSearch                  bool `json:"allow_index_search"`
			Release                           bool `json:"release"`
			Delete                            bool `json:"delete"`
			DeletePublicSearchFolder          bool `json:"delete_public_search_folder"`
			CreatePublicSearchFolder          bool `json:"create_public_search_folder"`
			ViewPublicFolder                  bool `json:"view_public_folder"`
			EditPreviousVersions              bool `json:"edit_previous_versions"`
			DisplayPublicDocuments            bool `json:"display_public_documents"`
			ReadOnly                          bool `json:"read_only"`
			ViewPublicSearchFolder            bool `json:"view_public_search_folder"`
			CreatePublicFolder                bool `json:"create_public_folder"`
			Import                            bool `json:"import"`
			ShareDocument                     bool `json:"share_document"`
		} `json:"m1_bits"`
		M2     int `json:"m2"`
		M2Bits struct {
			DocumentView   bool `json:"document_view"`
			External       bool `json:"external"`
			UseAdminTool   bool `json:"use_admin_tool"`
			UseImportTool  bool `json:"use_import_tool"`
			UseMonitorTool bool `json:"use_monitor_tool"`
		} `json:"m2_bits"`
		M3     int `json:"m3"`
		M3Bits struct {
			AuthorWorkspace        bool `json:"author_workspace"`
			BrowseWorkspace        bool `json:"browse_workspace"`
			CreateTemplate         bool `json:"create_template"`
			CustomizePersonalViews bool `json:"customize_personal_views"`
			CustomizePublicViews   bool `json:"customize_public_views"`
			DeleteWorkspace        bool `json:"delete_workspace"`
			SearchWorkspace        bool `json:"search_workspace"`
			ShareWorkspace         bool `json:"share_workspace"`
		} `json:"m3_bits"`
		M4Bits struct {
			AdminCustom              bool `json:"admin_custom"`
			AdminTier1               bool `json:"admin_tier_1"`
			AdminTier2               bool `json:"admin_tier_2"`
			ContentAssistance        bool `json:"content_assistance"`
			CustomMetadataManagement bool `json:"custom_metadata_management"`
			DmsJobOperations         bool `json:"dms_job_operations"`
			GovernSecurityManagement bool `json:"govern_security_management"`
			Report                   bool `json:"report"`
			ReportManagement         bool `json:"report_management"`
			RoleManagement           bool `json:"role_management"`
			SystemManagement         bool `json:"system_management"`
			SystemMetadataManagement bool `json:"system_metadata_management"`
			SystemJobOperations      bool `json:"system_job_operations"`
			Tier                     int  `json:"tier"`
			TrusteeAssistance        bool `json:"trustee_assistance"`
			TrusteeManagement        bool `json:"trustee_management"`
		} `json:"m4_bits"`
	} `json:"data"`
}

func GetRoleDetails(ctx context.Context, client *client.Client, alias string, libraryId string)(GetRoleDetailsOutput, error){
	var data GetRoleDetailsOutput
	endpoint := client.BaseUrl.JoinPath("work","api","v2","customers",client.TokenCfg.CustomerId, "libraries", libraryId, "roles", alias)

	req := internal.HttpRequestBuilder{}
	req.Context(ctx).Url(*endpoint).Method(internal.Get).ToJson(&data)

	if err := client.Req(req); err != nil {
		return data, err
	}
	return data, nil
}