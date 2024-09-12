# HTML Parser
## Description
Create a simple HTML parser that can parse the tags and attributes of all the HTML files located in `./html` directory and convert them into MD format. The parser should be able to handle the following tags:
- `h1`, `h2`, `h3`, `h4`, `h5`, `h6`
- `p`
- `a`
- `img`
- `ul`, `ol`, `li`

The parser should be able to handle the following attributes:
- `href` for `a` tag
- `src` for `img` tag

The parser should be able to create a valid MD file with the same name as the HTML file in the `./md` directory.

## Example
### Input
`./html/home.html`
```html
<html>
    <head>
        <title>Home</title>
    </head>
    <body>
        <h1>Home</h1>
        <h2>Home Heading 2</h2>
        <p>Home Paragraph</p>
        <p><a href="home.html">Home</a></p>
        <p><a href="about.html">About</a></p>
        <p><a href="career.html">Career</a></p>
    </body>
</html>
```

### Output
`./md/home.md`
```markdown
# Home
## Home Heading 2
Home Paragraph
[Home](home.md)
[About](about.md)
[Career](career.md)
```

## Note
This parser does not require any input from the user. When the program executes, it simply looks for the all the files in the `./html` directory and converts them into MD format in the `./md` directory.