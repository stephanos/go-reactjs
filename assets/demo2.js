/**
 * Source: http://facebook.github.io/react/tips/expose-component-functions.html
 */
var React = self.React;

var HelloWorld = React.createClass({displayName: 'HelloWorld',
  render: function() {
    return (
      React.DOM.p(null,
        "Hello, ", React.DOM.input({type: "text", placeholder: "Your name here"}), "!"
      )
    );
  }
});