// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by pipeline:
//     terraform
// Using jennies:
//     TerraformDataSourceJenny
//     LatestJenny
//
// Run 'go generate ./' from repository root to regenerate.

package datasources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/grafana/terraform-provider-grafana-framework/internal/common"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &playlistDataSource{}
	_ datasource.DataSourceWithConfigure = &playlistDataSource{}
)

func NewplaylistDataSource() datasource.DataSource {
	return &playlistDataSource{}
}

// playlistDataSource defines the data source implementation.
type playlistDataSource struct {
	client *common.Client
}

// playlistDataSourceModel describes the data source data model.
type playlistDataSourceModel struct {
	Uid      types.String `tfsdk:"uid", json:"uid"`
	Name     types.String `tfsdk:"name", json:"name"`
	Interval types.String `tfsdk:"interval", json:"interval"`
	Items    types.List   `tfsdk:"items", json:"items"`
	ToJSON   types.String `tfsdk:"to_json"`
}

func (d *playlistDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_playlist"
}

func (d *playlistDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "TODO description",

		Attributes: map[string]schema.Attribute{
			"uid": schema.StringAttribute{
				MarkdownDescription: `Unique playlist identifier. Generated on creation, either by the
creator of the playlist of by the application.`,
				Computed: false,
				Optional: false,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the playlist.`,
				Computed:            false,
				Optional:            false,
			},
			"interval": schema.StringAttribute{
				MarkdownDescription: `Interval sets the time between switching views in a playlist.
FIXME: Is this based on a standardized format or what options are available? Can datemath be used?`,
				Computed: false,
				Optional: false,
			},
			"items": schema.ListAttribute{
				MarkdownDescription: `The ordered list of items that the playlist will iterate over.
FIXME! This should not be optional, but changing it makes the godegen awkward`,
				Computed: false,
				Optional: true,
			},
			"to_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "This datasource rendered as JSON",
			},
		},
	}
}

func (d *playlistDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*common.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *common.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *playlistDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data playlistDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	JSONConfig, err := json.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("JSON marshalling error", err.Error())
		return
	}

	// Not sure about that
	data.ToJSON = types.StringValue(string(JSONConfig))

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
