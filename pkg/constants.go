//
// Copyright 2023-present Sean Ottey. All rights reserved.
// Use of this source code is governed by MIT license
// that can be found in the LICENSE file.
//

package pkg

import "github.com/AlecAivazis/survey/v2"

// Version is current version of application.
const Version = "v0.1.5"

var (
	// Custom configuration for survey icons and colors.
	// See [https://github.com/mgutz/ansi#style-format] for details.
	SurveyIconsConfig = func(icons *survey.IconSet) {
		icons.Question.Format = "cyan"
		icons.Question.Text = "[?]"
		icons.Help.Format = "blue"
		icons.Help.Text = "Help ->"
		icons.Error.Format = "yellow"
		icons.Error.Text = "Warning ->"
	}
)
