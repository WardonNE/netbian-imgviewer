package gui

import (
	"modules/app"

	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

func GetImageView() ImageView {
	imageview := app.MainWindowConf.Children.BodyComposite.ImageView
	maxsize := Size{
		imageview.MaxSize.Width,
		imageview.MaxSize.Height,
	}
	minsize := Size{
		imageview.MinSize.Width,
		imageview.MinSize.Height,
	}
	image := imageview.Image
	margin := imageview.Margin
	background := SolidColorBrush{
		Color: walk.RGB(imageview.Background.R, imageview.Background.G, imageview.Background.B),
	}
	return ImageView{
		AssignTo:   &mw.imageviewer,
		Image:      image,
		Background: background,
		Margin:     margin,
		Mode:       ImageViewModeShrink,
		MaxSize:    maxsize,
		MinSize:    minsize,
	}
}
