package server

const Index string = `
<html>
  <head>
    <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.7.2/css/all.css" integrity="sha384-fnmOCqbTlWIlj8LyTjo7mOUStjsKC4pOpQbqyi7RrhN7udi9RwhKkMHpvLbHG9Sr" crossorigin="anonymous">
    <style>
      body {
        margin: 0;
        padding: 0;
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue", sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
      }

      html, body, div, span, applet, object, iframe,
      h1, h2, h3, h4, h5, h6, p, blockquote, pre,
      a, abbr, acronym, address, big, cite, code,
      del, dfn, em, img, ins, kbd, q, s, samp,
      small, strike, strong, sub, sup, tt, var,
      b, u, i, center,
      dl, dt, dd, ol, ul, li,
      fieldset, form, label, legend,
      table, caption, tbody, tfoot, thead, tr, th, td,
      article, aside, canvas, details, embed,
      figure, figcaption, footer, header, hgroup,
      menu, nav, output, ruby, section, summary,
      time, mark, audio, video {
        margin: 0;
        padding: 0;
        border: 0;
        font-size: 100%;
        vertical-align: baseline;
      }
      /* HTML5 display-role reset for older browsers */
      article, aside, details, figcaption, figure,
      footer, header, hgroup, menu, nav, section {
        display: block;
      }
      body {
        line-height: 1;
      }
      ol, ul {
        list-style: none;
      }
      blockquote, q {
        quotes: none;
      }
      blockquote:before, blockquote:after,
      q:before, q:after {
        content: '';
        content: none;
      }
      table {
        border-collapse: collapse;
        border-spacing: 0;
      }

      code {
        font-family: source-code-pro, Menlo, Monaco, Consolas, "Courier New", monospace;
      }
      
      a {
        color: #fcfcfc;
      }

      a:visited {
        color: #fcfcfc;
      }

      .container {
        height: 100%;
        width: 100%;

        display: flex;
        justify-content: center;

        /* background-image: linear-gradient(#a1c4fd, #c2e9fb); */
        background-image: linear-gradient(#111111, #131313);
      }

      .icon {
        margin-right: 0.5rem;
      }

      .title {
        color: #fcfcfc;
        text-align: center;
        font-size: 50px;
        font-weight: 800;
        margin-bottom: 2rem;

        width: 100%;
        display: flex;
        justify-content: space-between;
      }

      .description {
        color: #fcfcfc;
        margin-bottom: 2rem;
        line-height: 120%;
      }

      .statusBox {
        width: 100%;
        line-height: 150%;
        margin: 0 0 2rem 0;
      }

      .success {
        color: #4CAF50;
      }

      .failure {
        color: #d62243;
      }

      .formContainer {
        height: 100%;
        
        width: 100%;
        margin: 0 2rem;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
      }

        @media (min-width: 600px) {
          .formContainer {
            width: 28rem;
            margin: 0;
          }
	      }

      .form {
        display: flex;
        flex-direction: column;
        width: 100%;
      }

      .form p {
        color: #fcfcfc;
        margin-bottom: 0.5rem;
      }

      .form input {
        margin-bottom: 1.25rem;
      }

      input[type=text] {
        padding: 0 5px;
        height: 36px;
        position: relative;
        left: 0;
        outline: none;
        border-radius: 6px;
        
        color: #fcfcfc;
        background-color: #131313;
        font-size: 14px;
      }

      input[type=text]:focus {
        border: 2px solid #1a93ce;
      }

      input[type=submit] {
        margin: 1rem auto;
        width: 6rem;
        background: transparent;
        border: 2px solid #0099CC;
        border-radius: 6px; 
        
        background-color: #131313; 
        border: 2px solid #fcfcfc;

        color: white;

        padding: 6px 16px;
        text-align: center;
        
        display: inline-block;
        font-size: 16px;

        -webkit-transition-duration: 0.25s; /* Safari */
        transition-duration: 0.25s;
        cursor: pointer;
        text-decoration: none;
      }

      input[type=submit]:hover {
        background-color: #fcfcfc;
        color: 131313;
      }

      footer {
        position: absolute;
        bottom: 0;
        width: 100%;
        
        margin-bottom: 1rem;

        display: flex;
        justify-content: center;

        color: #fcfcfc;
        font-size: 80%;
      }

      .footerContent {
        display: flex;
      }

      .footerContent p:not(:last-child) {
        margin-right: 1rem;
      }

      .source {
        position: absolute;
        top: 0;
        right: 0;

        margin: 1rem 1.5rem;
        
        color: #fcfcfc;
        font-size: 80%;
      }

      .source a {
        text-decoration: none;
      }

    </style>
  </head>
  <body>
    <div class="source">
      <a href="https://github.com/c-o-l-o-r/watchtower">source</a>
    </div>
    <div class="container">
      <div class="formContainer">
        <div class="title">
          <h1>WATCHTOWER</h1>
          <i class="far fa-building"></i>
        </div>
        <div class="description">
          Receive notifications of malicious activity on your Plasma chain via 
          text message and our <a href="https://twitter.com/PlasmaMonitor">Twitter bot</a>.
        </div>
        <script type = "text/javascript">
          function queryStringToObject(queryString) {
            const query = {};
            if (queryString === '') {
              return query;
            }
            const pairs = (queryString[0] === '?'
              ? queryString.substr(1)
              : queryString
            ).split('&');
            for (let i = 0; i < pairs.length; i++) {
              const pair = pairs[i].split('=');
              query[decodeURIComponent(pair[0])] = decodeURIComponent(pair[1] || '');
            }
            return query;
          }

          var urlParams = queryStringToObject(location.search);
          console.log(urlParams.success)

          if (urlParams.success === 'true') {
            document.write('<div class="statusBox success"><i class="icon fas fa-lock"></i>' + urlParams.address + '</div>')
          }

          if (urlParams.success === 'false') {
            document.write('<div class="statusBox failure"><i class="icon fas fa-lock-open"></i>Something went wrong, please try again.</div>')
          }
       
        </script>

        <form class="form" action="/watchtower/" method="post">
          <p>ethereum address</p>
          <input type="text" name="address">
          <p>phone number</p>
          <input type="text" name="phone">
          <input type="submit" value="watch">
        </form>
      </div>
      <footer>
        <div class="footerContent">
          <!-- <p>built @ eth denver</p> -->
        </div>
      </footer>
    </div>
  </body>
</html>
`
