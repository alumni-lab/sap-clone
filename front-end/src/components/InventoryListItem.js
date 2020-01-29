import React from "react";
// import "components/InventoryListItem.scss";


export default function InventoryListItem(props) {
  return (
    <li onClick={() => props.select}>
      <h2 >{props.name}</h2>
      <h3 >{props.items}</h3>
    </li>
  );
}