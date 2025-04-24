package impl

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/agenttool"
	"github.com/iter-x/iter-x/internal/data/impl/build"
)

type toolRepo struct {
	*data.Tx
}

// NewTool creates a new tool repository
func NewTool(d *data.Data) repository.ToolRepo {
	return &toolRepo{
		Tx: d.Tx,
	}
}

func (r *toolRepo) ToEntity(po *ent.AgentTool) *do.Tool {
	if po == nil {
		return nil
	}

	return &do.Tool{
		ID:          po.ID,
		CreatedAt:   po.CreatedAt,
		UpdatedAt:   po.UpdatedAt,
		Name:        po.Name,
		Type:        po.Type,
		Description: po.Description,
		BaseURL:     po.BaseURL,
		APIKey:      po.APIKey,
		Model:       po.Model,
		Function:    build.ConvertFunctionSpec(po.FunctionSpec),
		Enabled:     true,
	}
}

func (r *toolRepo) ToEntities(pos []*ent.AgentTool) []*do.Tool {
	if pos == nil {
		return nil
	}

	tools := make([]*do.Tool, 0, len(pos))
	for _, po := range pos {
		tools = append(tools, r.ToEntity(po))
	}
	return tools
}

// GetTool returns a tool by name
func (r *toolRepo) GetTool(ctx context.Context, name string) (*do.Tool, error) {
	tool, err := r.Tx.GetTx(ctx).AgentTool.Query().
		Where(agenttool.NameEQ(name)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return r.ToEntity(tool), nil
}

// ListTools returns all tools
func (r *toolRepo) ListTools(ctx context.Context) ([]*do.Tool, error) {
	tools, err := r.Tx.GetTx(ctx).AgentTool.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return r.ToEntities(tools), nil
}

// CreateTool creates a new tool
func (r *toolRepo) CreateTool(ctx context.Context, tool *do.Tool) error {
	_, err := r.Tx.GetTx(ctx).AgentTool.Create().
		SetName(tool.Name).
		SetType(tool.Type).
		SetNillableDescription(&tool.Description).
		SetNillableBaseURL(&tool.BaseURL).
		SetNillableAPIKey(&tool.APIKey).
		SetNillableModel(&tool.Model).
		SetFunctionSpec(build.ConvertFunctionConfig(tool.Function)).
		Save(ctx)
	return err
}

// UpdateTool updates an existing tool
func (r *toolRepo) UpdateTool(ctx context.Context, tool *do.Tool) error {
	return r.Tx.GetTx(ctx).AgentTool.Update().
		Where(agenttool.NameEQ(tool.Name)).
		SetType(tool.Type).
		SetNillableDescription(&tool.Description).
		SetNillableBaseURL(&tool.BaseURL).
		SetNillableAPIKey(&tool.APIKey).
		SetNillableModel(&tool.Model).
		SetFunctionSpec(build.ConvertFunctionConfig(tool.Function)).
		Exec(ctx)
}

// DeleteTool deletes a tool
func (r *toolRepo) DeleteTool(ctx context.Context, name string) error {
	_, err := r.Tx.GetTx(ctx).AgentTool.Delete().
		Where(agenttool.NameEQ(name)).
		Exec(ctx)
	return err
}
