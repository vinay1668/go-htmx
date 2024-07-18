// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func chat() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Chat Box Container --><div class=\"flex flex-col bg-[#0F172A] shadow-lg rounded-lg w-full max-w-md h-full max-h-96 overflow-hidden\"><!-- Chat messages --><div class=\"flex-1 p-4 overflow-y-auto\" id=\"chat-box\"><div class=\"mb-4\"><div class=\"bg-blue-500 text-white p-2 rounded-lg inline-block\">Hello! How can I help you?</div></div><div class=\"mb-4 text-right\"><div class=\"bg-green-500 text-white p-2 rounded-lg inline-block\">I need some information.</div></div><!-- More messages here --></div><!-- Input box --><div class=\"border-t border-gray-700 p-4\"><form id=\"chat-form\" class=\"flex\"><input type=\"text\" id=\"chat-input\" class=\"flex-1 border border-gray-600 rounded-lg p-2 focus:outline-none focus:ring-2 focus:ring-blue-600 bg-[#1E293B] text-white placeholder-gray-400\" placeholder=\"Type your message...\"> <button type=\"submit\" class=\"ml-2 bg-blue-600 text-white rounded-lg px-4 py-2 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-600\">Send</button></form></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}