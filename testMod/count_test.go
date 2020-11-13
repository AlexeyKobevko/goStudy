package count

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestCount(t *testing.T) {
	s := "qweqwrwsf"
	//e := 3
	//if c := Count(s, 'w'); c != e {
	//	t.Fatalf("Bad count for %s: got %d expected %d", s, c, e)
	//}
	require.Equal(t, Count(s, 'w'), 3, "counting 'w' in"+s)
	require.Equal(t, Count(s, 'q'), 2, "counting 'q' in"+s)
	require.Equal(t, Count(s, 'f'), 1, "counting 'f' in"+s)
}