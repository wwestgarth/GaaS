package primitives

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wwestgarth/remote-geometry/vector"
)

func TestCreateSheetCircle(t *testing.T) {

	sheet := CreateSheetCircle(10, vector.NewZeroVector(), vector.NewVector(0, 0, 1))
	require.NotNil(t, sheet)
	require.Nil(t, sheet.Validate())

	sheet = CreateSheetCircle(-10, vector.NewZeroVector(), vector.NewVector(0, 0, 1))
	require.NotNil(t, sheet)
	require.NotNil(t, sheet.Validate())
}
