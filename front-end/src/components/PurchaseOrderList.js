import React from "react";
import ReactDOM from "react-dom";
import PurchaseOrderItem from "../components/PurchaseOrderItem.js";

export default function PurchaseOrderList(props) {
  const POs = props.POs.map(PO => {
    return (
      <PurchaseOrderItem
        key={PO.id}
        POid={PO.POid} 
        customerName={PO.customerName} 
        amount={PO.amount} 
      />
    );
  });

  return POs;
}
