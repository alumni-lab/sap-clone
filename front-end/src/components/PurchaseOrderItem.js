import React from "react";
import "../components/InventoryListItem.scss";

// import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

var classnames = require('classnames');

export default function PurchaseOrderItem(props) {
  return (
    <TableRow key={props.key}>
      <TableCell>
        {props.POid}
      </TableCell>
      <TableCell>
        {props.customerName}
      </TableCell>
      <TableCell>
        {props.amount}
      </TableCell>
    </TableRow>
  );
}
