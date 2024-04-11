// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package index

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Base(title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" data-theme=\"dark\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/index/index.templ`, Line: 10, Col: 18}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><!-- Tailwind CSS --><link rel=\"stylesheet\" href=\"./static/css/tailwind.css\"><link rel=\"stylesheet\" href=\"https://fonts.cdnfonts.com/css/roboto\"><link href=\"https://fonts.cdnfonts.com/css/ibm-plex-sans\" rel=\"stylesheet\"><style>\n        @import url('https://fonts.cdnfonts.com/css/ibm-plex-sans');\n    </style><!-- HTMX --><script src=\"https://cdn.jsdelivr.net/npm/htmx.org/dist/htmx.min.js\"></script></head><body></body><body style=\"background-color:var(--subtle);\"><div class=\"flex justify-center items-center h-screen\"><div class=\"w-full max-w-screen-md px-10\"><div class=\"move-cursor\"><span class=\"font-robotoBlack text-xl\">weiting@weitingwork ~ $ ls</span></div><section class=\"font-ibmPlexSans text-base px-2 py-2\"><ul class=\"list-none\"><li><a href=\"/\">Portfoilo</a></li><li><a href=\"/\">Github</a></li><li><a href=\"/\">Linkedin</a></li><li><a href=\"/\">Resume</a></li><li><a href=\"/\">...</a></li></ul></section><section class=\"font-robotoMedium text-base px-2 py-2\"><p>I'm a developer, engineer, and architect who passionates about coding, robotics, and design.<br>As a software developer, I am working on backend development, desktop application, computer graphics and computer system. As a architect, I am focusing on computational design and robotic fabrication. In my free time, I am building my muscle at gym.</p></section></div></div></body>></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}