import React from "react";
import "../components/InventoryListItem.scss";

var classnames = require('classnames');

export default function PurchaseOrderItem(props) {
  return (
    <li className = "inventory-item">
      <h3 >{props.POid}</h3>
      <h2 >{props.customerName}</h2>
      <h3 >{props.amount}</h3>
    </li>
  );
}
