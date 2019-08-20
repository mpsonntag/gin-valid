package templates

// TODO: Stable access order of Hooks map
const RepoList = `
{{define "content"}}

<div class="explore repositories">
	<div class="ui container repository list">
		{{range .}}
			{{$repopath := .FullName}}
			<div class="item">
				<div class="ui grid">
					<div class="ui two wide column middle aligned center">
						<i class="mega-octicon octicon-repo"></i>
					</div>
					<div class="ui fourteen wide column">
						<div class="ui header">
							<a class="name" href="/repos/{{$repopath}}/hooks">{{$repopath}}</a>
							<div class="ui links text normal small">
								<span>Active validators</span>
						{{range $hookname, $hook := .Hooks}}
							{{if eq $hook.State 0}}
								<span> | {{$hookname | ToUpper}}: <a href="/results/{{$hookname | ToLower}}/{{$repopath}}">results</a> </span>
							{{end}}
						{{end}}
							</div>
						</div>
						<p class="has-emoji">{{.Description}}</p>
						<a href="{{.HTMLURL}}">Repository on GIN</a> | <a href="{{.HTMLURL}}/settings/hooks">Repository hooks</a>
					</div>
				</div>
			</div>
		{{end}}
	</div>
</div>
	{{end}}
`
const repoPage = `
{{define "content"}}
	<div class="repository file list">
		<div><b>{{.FullName}}</b></div>
		<div><b>Available validators</b>:<br>
		{{range $hookname, $hook := .Hooks}}
			{{$hookname}}
			{{if eq $hook.State 0}}
				[Enabled] <a href="/repos/{{$.FullName}}/{{$hook.ID}}/disable">disable</a>
			{{else}}
					[Disabled] <a href="/repos/{{$.FullName}}/{{$hookname}}/enable">enable</a>
			{{end}}
			<br>
		{{end}}
		</div>
		<div>{{.Description}} {{.Website}}</div>
		<hr>
	</div>
{{end}}
`

const RepoPage = `
{{define "content"}}
	<div class="repository file list">
		<div class="header-wrapper">
			<div class="ui container">
				<div class="ui vertically padded grid head">
					<div class="column">
						<div class="ui header">
							<div class="ui huge breadcrumb">
								<i class="mega-octicon octicon-repo"></i>
								{{.FullName}}
							</div>
						</div>
					</div>
				</div>
			</div>
			<div class="ui tabs container">
			</div>
			<div class="ui tabs divider"></div>
		</div>
		<div class="ui container">
			<p id="repo-desc">
			<span class="description has-emoji">{{.Description}}</span>
			<a class="link" href=""></a>
			</p>
			<table id="repo-files-table" class="ui unstackable fixed single line table">
				<tbody>
					{{range $hookname, $hook := .Hooks}}
						<tr>
							<td class="name text bold four wide"><a href="">{{$hookname | ToUpper}}</a></td>
							{{if eq $hook.State 0}}
								<td class="name nine wide"><a href="">RESULTS</a></td>
								<td class="name three wide"><a href="/repos/{{$.FullName}}/{{$hook.ID}}/disable">DEACTIVATE</a></td>
							{{else}}
								<td class="name nine wide">N/A</td>
								<td class="name three wide"><a href="/repos/{{$.FullName}}/{{$hookname}}/enable">ACTIVATE</a></td>
							{{end}}
						</tr>
					{{end}}
				</tbody>
			</table>
		</div>
	</div>
{{end}}
`
