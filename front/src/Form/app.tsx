import React from "react";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import TextField from "@material-ui/core/TextField";
import { makeStyles } from "@material-ui/core/styles";

import RadioButtonsGroupGender from "./Radio/RadioButtonsGroupGender";
import RadioButtonsGroupLegal from "./Radio/RadioButtonsGroupLegal";
import RadioButtonsGroupParent from "./Radio/RadioButtonsGroupParent";
import RadioButtonsGroupLicense from "./Radio/RadioButtonsGroupLicense";
import RadioButtonsGroupRaquette from "./Radio/RadioButtonsGroupRaquette";
import RadioButtonsGroupRepas from "./Radio/RadioButtonsGroupLicense";

import Textfields from "./Textfields/textfield";

import SimpleSelectSemaine from "./Select/selectSemaine";
import SimpleSelectFormule from "./Select/selectSemaine";

import DatePickers from "./Textfields/datepicker";
import TextfieldsParents from "./Textfields/TextfieldsParents";
import TextfieldsAddress from "./Textfields/TextfieldsAddress";
import TextfieldsCard from "./Textfields/TextfieldsCard";

import CheckboxLabels from "./Checkbox/conditions";

import Button from "@material-ui/core/Button";

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(8),
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: "100%", // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

/*
const useStyles = makeStyles((theme) => ({
    container: {
        display: 'flex',
        flexWrap: 'wrap',
    },
    textField: {
        marginLeft: theme.spacing(10),
        marginRight: theme.spacing(10),
        width: 200,
    },
})); */
// export default function SignUp() {
//     const classes = useStyles();

export default function AddressForm() {
  return (
    <form onSubmit={(e) => console.log(e)}>
      <Typography variant="h6" gutterBottom>
        STAGES DE TENNIS
      </Typography>
      <Textfields />
      <Grid container spacing={3}>
        <Grid item xs={12}>
          <DatePickers />
          <RadioButtonsGroupGender />
          <RadioButtonsGroupLicense />
          <SimpleSelectFormule />
          <SimpleSelectSemaine />
          <RadioButtonsGroupRaquette />
          <RadioButtonsGroupRepas />
          <TextfieldsParents />
          <RadioButtonsGroupParent />
          <TextfieldsAddress />
          <CheckboxLabels />
          <TextfieldsCard />
          <Button type="submit" variant="contained"> 
            S'inscrire
          </Button>
        </Grid>
      </Grid>
    </form>
  );
}
