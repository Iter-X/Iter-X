package impl

import (
	"context"

	"github.com/iter-x/iter-x/internal/biz/do"
	"github.com/iter-x/iter-x/internal/biz/repository"
	"github.com/iter-x/iter-x/internal/data"
	"github.com/iter-x/iter-x/internal/data/ent"
	"github.com/iter-x/iter-x/internal/data/ent/agent"
	"github.com/iter-x/iter-x/internal/data/ent/agentprompt"
	"github.com/iter-x/iter-x/internal/data/ent/agenttoolbinding"
	"github.com/iter-x/iter-x/internal/data/impl/build"
)

type agentRepo struct {
	*data.Tx
}

// NewAgent creates a new agent repository
func NewAgent(d *data.Data) repository.AgentRepo {
	return &agentRepo{
		Tx: d.Tx,
	}
}

func (r *agentRepo) ToEntity(po *ent.Agent) *do.Agent {
	if po == nil {
		return nil
	}
	return build.AgentRepositoryImplToEntity(po)
}

func (r *agentRepo) ToEntities(pos []*ent.Agent) []*do.Agent {
	if pos == nil {
		return nil
	}
	return build.AgentRepositoryImplToEntities(pos)
}

// GetAgent returns an agent by name
func (r *agentRepo) GetAgent(ctx context.Context, name string) (*do.Agent, error) {
	a, err := r.Tx.GetTx(ctx).Agent.Query().
		Where(agent.NameEQ(name)).
		WithPrompts().
		WithToolBindings(func(q *ent.AgentToolBindingQuery) {
			q.WithTool()
		}).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return r.ToEntity(a), nil
}

// ListAgents returns all agents
func (r *agentRepo) ListAgents(ctx context.Context) ([]*do.Agent, error) {
	agents, err := r.Tx.GetTx(ctx).Agent.Query().
		WithPrompts().
		WithToolBindings(func(q *ent.AgentToolBindingQuery) {
			q.WithTool()
		}).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return r.ToEntities(agents), nil
}

// CreateAgent creates a new agent
func (r *agentRepo) CreateAgent(ctx context.Context, agent *do.Agent) error {
	return r.Tx.WithTx(ctx, func(tx *ent.Tx) error {
		// Create agent
		a, err := tx.Agent.Create().
			SetName(agent.Name).
			SetEnabled(agent.Enabled).
			SetNillableDescription(&agent.Description).
			Save(ctx)
		if err != nil {
			return err
		}

		// Create prompt
		if agent.Prompt != nil {
			_, err = tx.AgentPrompt.Create().
				SetVersion(agent.Prompt.Version).
				SetSystem(agent.Prompt.System).
				SetUser(agent.Prompt.User).
				SetAgent(a).
				Save(ctx)
			if err != nil {
				return err
			}
		}

		// Create tools and bindings
		for _, tool := range agent.Tools {
			t, err := tx.AgentTool.Create().
				SetName(tool.Name).
				SetType(tool.Type).
				SetNillableDescription(&tool.Description).
				SetNillableBaseURL(&tool.BaseURL).
				SetNillableAPIKey(&tool.APIKey).
				SetNillableModel(&tool.Model).
				SetFunctionSpec(build.ConvertFunctionConfig(tool.Function)).
				Save(ctx)
			if err != nil {
				return err
			}

			// Create binding
			_, err = tx.AgentToolBinding.Create().
				SetAgent(a).
				SetTool(t).
				SetEnabled(tool.Enabled).
				Save(ctx)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// UpdateAgent updates an existing agent
func (r *agentRepo) UpdateAgent(ctx context.Context, agentDo *do.Agent) error {
	return r.Tx.WithTx(ctx, func(tx *ent.Tx) error {
		a, err := tx.Agent.Query().
			Where(agent.NameEQ(agentDo.Name)).
			Only(ctx)
		if err != nil {
			return err
		}

		err = tx.Agent.UpdateOne(a).
			SetEnabled(agentDo.Enabled).
			SetNillableDescription(&agentDo.Description).
			Exec(ctx)
		if err != nil {
			return err
		}

		// Update prompt if exists
		if agentDo.Prompt != nil {
			// Delete existing prompt
			_, err := tx.AgentPrompt.Delete().
				Where(agentprompt.HasAgentWith(agent.NameEQ(agentDo.Name))).
				Exec(ctx)
			if err != nil {
				return err
			}

			// Create new prompt
			_, err = tx.AgentPrompt.Create().
				SetVersion(agentDo.Prompt.Version).
				SetSystem(agentDo.Prompt.System).
				SetUser(agentDo.Prompt.User).
				SetAgent(a).
				Save(ctx)
			if err != nil {
				return err
			}
		}

		// Delete existing tools and bindings
		_, err = tx.AgentToolBinding.Delete().
			Where(agenttoolbinding.HasAgentWith(agent.NameEQ(agentDo.Name))).
			Exec(ctx)
		if err != nil {
			return err
		}

		// Create new tools and bindings
		for _, tool := range agentDo.Tools {
			t, err := tx.AgentTool.Create().
				SetName(tool.Name).
				SetType(tool.Type).
				SetNillableDescription(&tool.Description).
				SetNillableBaseURL(&tool.BaseURL).
				SetNillableAPIKey(&tool.APIKey).
				SetNillableModel(&tool.Model).
				SetFunctionSpec(build.ConvertFunctionConfig(tool.Function)).
				Save(ctx)
			if err != nil {
				return err
			}

			// Create binding
			_, err = tx.AgentToolBinding.Create().
				SetAgent(a).
				SetTool(t).
				SetEnabled(tool.Enabled).
				Save(ctx)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// DeleteAgent deletes an agent
func (r *agentRepo) DeleteAgent(ctx context.Context, name string) error {
	_, err := r.Tx.GetTx(ctx).Agent.Delete().
		Where(agent.NameEQ(name)).
		Exec(ctx)
	return err
}
