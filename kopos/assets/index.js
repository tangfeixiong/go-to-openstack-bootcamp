import "./stylesheets/base.scss";
import "material-design-lite";

import React from 'react';
import ReactDom from 'react-dom';

import Button from '#coms/Button';

ReactDom.render(
  <Button icon="search" title="MdlButton" color="primary" raised />, 
  document.getElementById("mount"));
