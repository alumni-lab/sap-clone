import React from "react";
// import "../components/InventoryListItem.scss";

// var classnames = require('classnames');

export default function Header(props) {

  return (
    <li>
      {props.username ? <h3>{props.username}</h3> : <h3>Login</h3>}
      <img src={props.logo} alt="logo"/>
    </li>
  );
}
