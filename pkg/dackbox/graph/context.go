package graph

import "context"

// Context executes the input function using the input graph provider with a context that provides a lazy generated
// graph.
func Context(ctx context.Context, provider Provider, fn func(hasGraphConext context.Context)) {
	inner := context.WithValue(ctx, graphContextKey{}, &graphContext{
		provider: provider,
	})
	defer discardGraph(inner)
	fn(inner)
}

// GetGraph returns the RGraph from the input context, or nil if the input context has no graph context.
func GetGraph(hasGraphContext context.Context) RGraph {
	gc, ok := hasGraphContext.Value(graphContextKey{}).(*graphContext)
	if !ok {
		return nil
	}
	if gc.graph == nil {
		gc.graph = gc.provider.NewGraphView()
	}
	return gc.graph
}

// Key for fetching graph instance from the contenxt.
type graphContextKey struct{}

type graphContext struct {
	provider Provider
	graph    DiscardableRGraph
}

func discardGraph(ctx context.Context) {
	gc, ok := ctx.Value(graphContextKey{}).(*graphContext)
	if !ok {
		return
	}
	if gc.graph == nil {
		return
	}
	gc.graph.Discard()
}
