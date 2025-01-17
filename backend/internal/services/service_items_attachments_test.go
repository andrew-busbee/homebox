package services

import (
	"context"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/hay-kot/homebox/backend/internal/repo"
	"github.com/stretchr/testify/assert"
)

func TestItemService_AddAttachment(t *testing.T) {
	temp := os.TempDir()

	svc := &ItemService{
		repo:     tRepos,
		filepath: temp,
	}

	loc, err := tSvc.Location.Create(context.Background(), tGroup.ID, repo.LocationCreate{
		Description: "test",
		Name:        "test",
	})
	assert.NoError(t, err)
	assert.NotNil(t, loc)

	itmC := repo.ItemCreate{
		Name:        fk.Str(10),
		Description: fk.Str(10),
		LocationID:  loc.ID,
	}

	itm, err := svc.Create(context.Background(), tGroup.ID, itmC)
	assert.NoError(t, err)
	assert.NotNil(t, itm)
	t.Cleanup(func() {
		err := svc.repo.Items.Delete(context.Background(), itm.ID)
		assert.NoError(t, err)
	})

	contents := fk.Str(1000)
	reader := strings.NewReader(contents)

	// Setup
	afterAttachment, err := svc.AttachmentAdd(tCtx, itm.ID, "testfile.txt", "attachment", reader)
	assert.NoError(t, err)
	assert.NotNil(t, afterAttachment)

	// Check that the file exists
	storedPath := afterAttachment.Attachments[0].Document.Path

	// {root}/{group}/{item}/{attachment}
	assert.Equal(t, path.Join(temp, "homebox", tGroup.ID.String(), "documents"), path.Dir(storedPath))

	// Check that the file contents are correct
	bts, err := os.ReadFile(storedPath)
	assert.NoError(t, err)
	assert.Equal(t, contents, string(bts))

}
