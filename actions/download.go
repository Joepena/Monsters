package actions

import "github.com/gobuffalo/buffalo"

func downloadHandler(c buffalo.Context) error {
	assetID := c.Param("assetID")
	mContext := c.(MonsterContext)

	return mContext.RenderFile(assetID)
}
