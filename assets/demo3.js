/**
 * Source: http://facebook.github.io/react/docs/tutorial.html
 */

var React = self.React;

var CommentBox = React.createClass({displayName: 'CommentBox',
  render: function () {
    var commentList = React.createElement(CommentList, 
      this.props
    );
    return (
      React.createElement(
        "div",
        {},
        React.DOM.h1(null, "Comments"),
        commentList
      )
    );
  }
});

var CommentList = React.createClass({displayName: 'CommentList',
  render: function () {
    var commentNodes = this.props.data.map(function (comment) {
      return React.createElement(
        Comment, 
        {key: comment.id, author: comment.author},
        comment.text
      );
    });
    return (
      React.DOM.div({className: "commentList"},
        commentNodes
      )
    );
  }
});

var Comment = React.createClass({displayName: 'Comment',
  render: function () {
    return (
      React.DOM.div({className: "comment"},
        React.DOM.h2({className: "commentAuthor"}, this.props.author),
        this.props.children
      )
    );
  }
});