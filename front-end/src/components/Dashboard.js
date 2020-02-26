import React from "react";
import ReactDOM from "react-dom";
import PurchaseOrderList from "../components/PurchaseOrderList.js";
import InventoryItemList from "../components/InventoryItemList.js";


export default function Dashboard(props) {
  const [view, setView] = useState(props.view || "");



  return (
    // searchbar
    // button
    // InventoryList
  );
  // const POs = props.POs.map(PO => {
  //   return (
  //     <PurchaseOrderItem
  //       key={PO.id}
  //       POid={PO.POid} 
  //       customerName={PO.customerName} 
  //       amount={PO.amount} 
  //     />
  //   );
  // });

  // return POs;
}
