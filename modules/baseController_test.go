package modules

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventUpdate(t *testing.T) {

	mapUpdateProperties := makeMapUpdateProperties(
		`{"Name":"Test!!!"}`,
		`{"Name":"Test!!! Change"}`,
	)

	fmt.Println(mapUpdateProperties["Name"])

	assert.Equal(t, "Test!!!", mapUpdateProperties["Name"].NewValue)
	assert.Equal(t, "Test!!! Change", mapUpdateProperties["Name"].OldValue)
}

func TestCheckDiff(t *testing.T) {
	//타겟인 Name이 이전과 지금값이 같으므로 무시
	check := checkDiff(map[string]UpdateProperty{
		"Name": UpdateProperty{
			NewValue: "new",
			OldValue: "new",
		}},
		NotificationType{
			Name:     "test",
			DiffMode: true,
			DiffKey:  "Name",
		})
	assert.False(t, check)

	//타겟인 Name이 이전과 지금값이 다르므로 이벤트 포함
	check = checkDiff(map[string]UpdateProperty{
		"Name": UpdateProperty{
			NewValue: "new",
			OldValue: "old",
		}},
		NotificationType{
			Name:     "test",
			DiffMode: true,
			DiffKey:  "Name",
		})

	assert.True(t, check)
}
