import React from "react";
import ReactDOM from "react-dom";
import InventoryListItem from "../components/InventoryListItem.js";

export default function InventoryList(props) {
  const items = props.items.map(item => {
    return (
      <InventoryListItem
        key={item.id}
        name={item.name} 
        quantity={item.quantity} 
        description={item.description} 
        type={item.type} 
        vendor={item.vendor} 
      />
    );
  });

  return items;
}
