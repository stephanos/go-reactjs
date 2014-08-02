/**
 * Source: https://github.com/robertkrimen/otto/issues/67
 */
var React = self.React;

var CommentBox = React.createClass({
  render: function () {
    return (
      React.DOM.div({
        className: 'commentBox',
        children: 'Hello, world! I am a CommentBox.'
      })
    );
  }
});