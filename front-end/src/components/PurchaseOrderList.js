import React from "react";
import ReactDOM from "react-dom";
import PurchaseOrderItem from "../components/PurchaseOrderItem.js";

import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

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

  return (
    <TableContainer component={Paper}>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>
              Order ID
            </TableCell>
            <TableCell>
              Customer Name
            </TableCell>
            <TableCell>
              Amount
            </TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {POs}
        </TableBody>

      </Table>
    </TableContainer>
  );
}
