import React from 'react';
import classnames from 'classnames';
import _ from 'lodash';

function mdlUpgrade(Component) {
  const MDLComponent = React.createClass({
    componentDidMount() {
      componentHandler.upgradeDom();
    },
    componentDidUpdate() {
      componentHandler.upgradeDom();
    },
    render() {
      return <Component {...this.props} />
    },
  });
  return MDLComponent;
}

const Icon = (props) => {
    return (<i className='material-icons' style={props.style}>{props.name}</i>);
};

const ClassedButton = (props) => {
  let btnProps = {
    className: props.className,
    to: props.to,
    style: props.style,
    onClick: (evt) => props.onClick && props.onClick(evt),
  };
  if (props.id) {
    btnProps.id = props.id;
  }
  let children = [];
  let toReverse = false;
  if (props.icon) {
    const alignRight = props.iconAlign === 'right';
    toReverse = alignRight;
    let iconStyle = _.assign({}, props.iconStyle);
    if (props.title) {
      alignRight ? iconStyle['marginLeft'] = 8 : iconStyle['marginRight'] = 8;
    }
    children.push(<Icon style={iconStyle} name={props.icon} key="icon" />);
  }
  if (props.title) {
    children.push(<span key='title'>{props.title}</span>);
  }
  if (toReverse) {
    children.reverse();
  }
  return props.to ? <Link {...btnProps}>{children}</Link> : <button {...btnProps}>{children}</button>;
};

const Button = (props) => {
  const claz = classnames('mdl-button mdl-js-button mdl-js-ripple-effect', {
    'mdl-button--colored': props.color === 'colored',               
    'mdl-button--primary': props.color === 'primary',
    'mdl-button--accent': props.color === 'accent',
    'mdl-button--raised': props.raised === true,
  });
  return <ClassedButton className={claz} {...props} />;
};

export default mdlUpgrade(Button);
