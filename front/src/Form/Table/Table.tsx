import React from 'react';
import { withStyles, Theme, createStyles, makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import clsx from 'clsx';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import ListItemText from '@material-ui/core/ListItemText';
import Select from '@material-ui/core/Select';
import Checkbox from '@material-ui/core/Checkbox';
import Chip from '@material-ui/core/Chip';
import {stageData} from '../../data';

async function getDatas (Semaine:string) {
    return fetch('http://127.0.0.1:8080/semaine?semaine='+Semaine, {
  method: 'GET',
    }).then((Response)=> {
      console.log(Response)
      return Response.json
    })
  .catch((err)=>console.error(err))
};

const handleChange = (event: React.ChangeEvent<{ name?: string; value: unknown }>, Semaine: any, setSemaine:any) => {
  /*async function toto() {
    setSemaine(await getDatas(event.target.value as string))
  }
  toto()*/
    setSemaine(stageData)
  ;
}
function NativeSelects (props: any) {
  
  return (
    <div>
      <FormControl style={{
      margin: 20,
      minWidth: 120,
    }}>
        <InputLabel htmlFor="semaine-stage">Semaine</InputLabel>
        <Select
          native
          value={props.Semaine&&props.Semaine.age}
          onChange={e => handleChange(e, props.Semaine, props.setSemaine)}
          inputProps={{
            name: 'semaine',
            id: 'semaine-stage',
          }}
        >
          <option aria-label="None" value="" />
          <option value={'SemaineA'}>Semaine A</option>
          <option value={'SemaineB'}>Semaine B</option>
          <option value={'SemaineC'}>Semaine C</option>
          <option value={'SemaineD'}>Semaine D</option>
          <option value={'SemaineE'}>Semaine E</option>
          <option value={'SemaineF'}>Semaine F</option>
          <option value={'SemaineG'}>Semaine G</option>
          <option value={'SemaineH'}>Semaine H</option>
        </Select>
      </FormControl> 
    </div>
  );
}

const StyledTableCell = withStyles((theme: Theme) =>
  createStyles({
    head: {
      backgroundColor: theme.palette.common.black,
      color: theme.palette.common.white,
    },
    body: {
      fontSize: 14,
    },
  }),
)(TableCell);

const StyledTableRow = withStyles((theme: Theme) =>
  createStyles({
    root: {
      '&:nth-of-type(odd)': {
        backgroundColor: theme.palette.action.hover,
      },
    },
  }),
)(TableRow);

function createData(nom: string, prenom: string, anniversaire: number, sexe: string, repas: string, formuler: string, niveau: string,  location: string) {
  return { nom, prenom, anniversaire, sexe, repas, formuler, niveau, location };
}

const rows = [
  createData('jean', 'michel', 67, 'H', 'picnic', 'tennis', 'debutant', 'oui'),

];

const useStyles = makeStyles({
  table: {
    minWidth: 700,
  },
});

export default function Tableau() {
  const classes = useStyles();

  const [Semaine, setSemaine] = React.useState<any>(undefined)
  console.log(Semaine)

  return (
    <TableContainer component={Paper}>
      <NativeSelects Semaine={Semaine} setSemaine={setSemaine}/>
      <Table className={classes.table} aria-label="customized table">
        <TableHead>
          <TableRow>
            <StyledTableCell>Nom</StyledTableCell>
            <StyledTableCell align="right">Prénom</StyledTableCell>
            <StyledTableCell align="right">Age</StyledTableCell>
            <StyledTableCell align="right">Sexe</StyledTableCell>
            <StyledTableCell align="right">Niveau</StyledTableCell>
            <StyledTableCell align="right">Formule</StyledTableCell>
            <StyledTableCell align="right">Repas</StyledTableCell>
            <StyledTableCell align="right">Raquette</StyledTableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {Semaine&&Semaine.tab.map((row:any) => (
            <StyledTableRow key={row.nom}>
              <StyledTableCell component="th" scope="row">
                {row.nom}
              </StyledTableCell>
                <StyledTableCell align="right">{row.prenom}</StyledTableCell>
                <StyledTableCell align="right">{row.anniversaire}</StyledTableCell>
                <StyledTableCell align="right">{row.sexe}</StyledTableCell>
                <StyledTableCell align="right">{row.niveau}</StyledTableCell>
                <StyledTableCell align="right">{row.formuler}</StyledTableCell>
                <StyledTableCell align="right">{row.repas}</StyledTableCell>
                <StyledTableCell align="right">{row.location}</StyledTableCell>
            </StyledTableRow>
          ))}
        </TableBody>
        {/* <Stages />  */}
      </Table>
    </TableContainer>
  );
}
/*
const Stages = () => {
  return (
    <>
    <div className ="places-stages" >Données des stages</div>
    </>
  );
}*/