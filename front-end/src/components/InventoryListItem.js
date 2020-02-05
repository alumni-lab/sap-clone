import React from "react";
import "../components/InventoryListItem.scss";

var classnames = require('classnames');


export default function InventoryListItem(props) {
  return (
    <li className = "inventory-item">
      <h2 >{props.name}</h2>
      <h3 >{props.quantity}</h3>
      <p >{props.description}</p>
      <p >{props.type}</p>
      <h3 >{props.vendor}</h3>
    </li>
  );
}