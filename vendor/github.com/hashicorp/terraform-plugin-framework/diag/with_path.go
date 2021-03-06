package diag

import (
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ DiagnosticWithPath = withPath{}

// withPath wraps a diagnostic with path information.
type withPath struct {
	Diagnostic

	path *tftypes.AttributePath
}

// Equal returns true if the other diagnostic is wholly equivalent.
func (d withPath) Equal(other Diagnostic) bool {
	o, ok := other.(withPath)

	if !ok {
		return false
	}

	if !d.Path().Equal(o.Path()) {
		return false
	}

	if d.Diagnostic == nil {
		return d.Diagnostic == o.Diagnostic
	}

	return d.Diagnostic.Equal(o.Diagnostic)
}

// Path returns the diagnostic path.
func (d withPath) Path() *tftypes.AttributePath {
	return d.path
}

// WithPath wraps a diagnostic with path information or overwrites the path.
func WithPath(path *tftypes.AttributePath, d Diagnostic) DiagnosticWithPath {
	wp, ok := d.(withPath)

	if !ok {
		return withPath{
			Diagnostic: d,
			path:       path,
		}
	}

	wp.path = path

	return wp
}
