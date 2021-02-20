import * as React from 'react';
import { Component } from 'react';

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


export default class NameForm extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        child: {
          prenom: '',
          nom: ''
        }
      };
    }
  
    handleChange = (event) => {
    }
  
    handleSubmit = (event) => {
      let body = {
          child:
           {
             age: 18,
             prenom: 'nicolas'}
           }

      console.log(body)
      fetch('http://127.0.0.1:8080/hello', {
          method: 'POST',
          body: body
        }).then(function(response) {
          console.log(response)
          return response;
        });
  
      event.preventDefault();
  }
  
    render() {
      return (
        <form onSubmit={this.handleSubmit}>

      <Typography variant="h6" gutterBottom>
        STAGES DE TENNIS
      </Typography>

      <label>
            Prénom gosse:
            <input type="text" value={this.state.value} name="child.prenom" onChange={this.handleChange} /><br/>
          </label>
          <label>
            Nom gosse:<input type="text" value={this.state.value} name="child.nom" onChange={this.handleChange} /><br/>
          </label>
          <input type="submit" value="Submit" />
        </form>
      );
    }
  }