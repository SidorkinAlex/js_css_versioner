package version_replace_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/SidorkinAlex/js_css_versioner/internal/models"
	"github.com/SidorkinAlex/js_css_versioner/internal/version_replace"
)

func TestVersionReplace_Execute(t *testing.T) {
	type args struct {
		source  io.Reader
		version int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "it should replace replace js & css extensions with new version",
			args: args{
				source: bytes.NewReader([]byte(`<html>
													<head>
														<script src="myscripts.js"></script>
														<script src='myscripts.js'></script>
														<script src='myscripts.js?v=1709881862'></script>
														<link rel="stylesheet" href="mystyle.css">
														<link rel="stylesheet" href='mystyle.css'>
														<link rel="stylesheet" href='mystyle.css?v=1709883557'>
													</head>
													<body>
														<div>
															<span>Hello world!</span>
														</div>
													<body>
												</html>`),
				),
				version: 1000001,
			},
			want: `<html>
						<head>
							<script src="myscripts.js?v=1000001"></script>
							<script src='myscripts.js?v=1000001'></script>
							<script src='myscripts.js?v=1000001'></script>
							<link rel="stylesheet" href="mystyle.css?v=1000001">
							<link rel="stylesheet" href='mystyle.css?v=1000001'>
							<link rel="stylesheet" href='mystyle.css?v=1000001'>
						</head>
						<body>
							<div>
								<span>Hello world!</span>
							</div>
						<body>
				</html>`,
		},
		{
			name: "it should do nothing when no js or css extensions",
			args: args{
				source: bytes.NewReader([]byte(`<html>
													<head\>
													<body>
														<div>
															<span>Hello world!</span>
														</div>
													<body>
												</html>`),
				),
				version: 1000001,
			},
			want: `<html>
						<head\>
						<body>
							<div>
								<span>Hello world!</span>
							</div>
						<body>
					</html>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpl, errN := version_replace.New(models.VersionedExtensionJS, models.VersionedExtensionCSS)
			if errN != nil {
				t.Errorf("VersionReplace.New() error = %v", errN)
			}
			destination := &bytes.Buffer{}
			if err := rpl.Execute(tt.args.source, destination, tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("VersionReplace.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			want := strings.ReplaceAll(tt.want, "\t", "")
			got := strings.ReplaceAll(destination.String(), "\t", "")

			if got != want {
				t.Errorf("VersionReplace.Execute() = %v, want %v", got, want)
			}
		})
	}
}
