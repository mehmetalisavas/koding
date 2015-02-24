package templates

var LoggedInHome = `
<!doctype html>
  <html lang="en">
  <head>
    {{template "header" . }}

    <link rel="stylesheet" href="/a/p/p/kd.css" />
    <link rel="stylesheet" href="/a/p/p/app.css" />
  </head>

  <body class='logged-in dark ide'>
    <!--[if IE]><script>(function(){window.location.href='/unsupported.html'})();</script><![endif]-->

    <script src="/a/p/p/thirdparty/pubnub.min.js"></script>
    <script src="/a/p/p/thirdparty/gapi.js"></script>
    <script src="/a/p/p/common.js"></script>
    <script src="/a/p/p/app.js"></script>

    <script>
      require('app')({
        config: {{.Runtime}},
        userId: {{.User.GetWithDefault "UserId" "null" }},
        userAccount: {{.User.GetWithDefault "Account" "null" }},
        userMachines: {{.User.GetWithDefault "Machines" "null"}},
        userWorkspaces: {{.User.GetWithDefault "Workspaces" "null"}},
        currentGroup: {{.User.GetWithDefault "Group" "null"}},
        envData: {{.User.GetWithDefault "EnvData" "{}"}},
        isLoggedInOnLoad: true,
        socialApiData: {{.User.GetWithDefault "SocialApiData" "null"}}
      });
    </script>

    <script>
      (function(d) {
        var config = {
          kitId: 'rbd0tum',
          scriptTimeout: 3000
        },
        h=d.documentElement,t=setTimeout(function(){h.className=h.className.replace(/\bwf-loading\b/g,"")+" wf-inactive";},config.scriptTimeout),tk=d.createElement("script"),f=false,s=d.getElementsByTagName("script")[0],a;h.className+=" wf-loading";tk.src='//use.typekit.net/'+config.kitId+'.js';tk.async=true;tk.onload=tk.onreadystatechange=function(){a=this.readyState;if(f||a&&a!="complete"&&a!="loaded")return;f=true;clearTimeout(t);try{Typekit.load(config)}catch(e){}};s.parentNode.insertBefore(tk,s)
      })(document);
    </script>

    {{template "analytics" }}

    {{if not .Impersonating }}
      <script type="text/javascript">
        (function () {
        var _user_id = '{{.User.GetWithDefault "Username" ""}}'; var _session_id = '{{.User.GetWithDefault "SessionId" ""}}'; var _sift = _sift || []; _sift.push(['_setAccount', 'f270274999']); _sift.push(['_setUserId', _user_id]); _sift.push(['_setSessionId', _session_id]); _sift.push(['_trackPageview']); (function() { function ls() { var e = document.createElement('script'); e.type = 'text/javascript'; e.async = true; e.src = ('https:' == document.location.protocol ? 'https://' : 'http://') + 'cdn.siftscience.com/s.js'; var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(e, s); } if (window.attachEvent) { window.attachEvent('onload', ls); } else { window.addEventListener('load', ls, false); } })();
        })();
      </script>
    {{end}}
</body>
</html>
`
