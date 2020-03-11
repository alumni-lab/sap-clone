import React, { useState } from "react";
import ReactDOM from "react-dom";
// import PurchaseOrderList from "../components/PurchaseOrderList.js";
import InventoryList from "./InventoryList.js";
import PurchaseOrderList from "./PurchaseOrderList.js";


export default function Dashboard(props) {
  // const [view, setView] = useState(props.view || "");
  const [view, setView] = useState("Inventory");

  return (
    <div>
      {view === "Inventory" ? (<InventoryList items={props.items} />) : (<PurchaseOrderList />)}
      <h1>hello</h1>
    </div>
  );

}

