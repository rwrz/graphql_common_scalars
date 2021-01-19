package graphql_common_scalars

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalDateTimeScalar(t time.Time) graphql.Marshaler {
	if t.IsZero() {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Format(time.RFC3339)))
	})
}

func UnmarshalDateTimeScalar(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		fmt.Printf(tmpStr)
		return time.Parse(time.RFC3339, tmpStr)
	}
	return time.Time{}, errors.New("time should be RFC3339 formatted string")
}
