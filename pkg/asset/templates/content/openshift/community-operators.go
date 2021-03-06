package openshift

import (
	"os"
	"path/filepath"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	communityOperatorsFilename = "community-operators.yaml"
)

var _ asset.WritableAsset = (*CommunityOperators)(nil)

// CommunityOperators is the operatorHub setting for OKD to use community operators only
type CommunityOperators struct {
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *CommunityOperators) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *CommunityOperators) Name() string {
	return "Community Operators setting"
}

// Generate generates the actual files by this asset
func (t *CommunityOperators) Generate(parents asset.Parents) error {
	data, err := content.GetOpenshiftTemplate(communityOperatorsFilename)
	if err != nil {
		return err
	}
	t.FileList = append(t.FileList, &asset.File{
		Filename: filepath.Join(content.TemplateDir, communityOperatorsFilename),
		Data:     []byte(data),
	})
	return nil
}

// Files returns the files generated by the asset.
func (t *CommunityOperators) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *CommunityOperators) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, communityOperatorsFilename))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = append(t.FileList, file)

	return true, nil
}
