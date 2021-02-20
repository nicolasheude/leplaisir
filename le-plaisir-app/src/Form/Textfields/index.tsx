import React from 'react';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import { makeStyles } from '@material-ui/core/styles';
import FormControl from '@material-ui/core/FormControl';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import RadioButtonsGroupGender from './Radio/radio';
import RadioButtonsGroupLicense from './Radio/radio';
import RadioButtonsGroupRaquette from './Radio/radio';
import RadioButtonsGroupRepas from './Radio/radio';
import SimpleSelectSemaine from './Form/Select/selectSemaine';
import SimpleSelectFormule from './Select/selectSemaine';
import DatePickers from './Textfields/selectSemaine';

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
}));

export default function AddressForm() {
    return (
        <React.Fragment>
            <Typography variant="h6" gutterBottom>
                STAGES DE TENNIS
      </Typography>
            <TextFields />
            <Grid container spacing={3}>
                <Grid item xs={12}>
                    <DatePickers />
                </Grid>
                <Grid item xs={12}>
                    <RadioButtonsGroupGender />
                </Grid>
                <Grid item xs={12}>
                    <RadioButtonsGroupLicense />
                </Grid>
                <Grid item xs={12}>
                    <SimpleSelectFormule />
                </Grid>
                <Grid item xs={12}>
                    <SimpleSelectSemaine />
                </Grid>
                <Grid item xs={12}>
                    <RadioButtonsGroupRaquette />
                </Grid>
                <Grid item xs={12}>
                    <RadioButtonsGroupRepas />
                </Grid>
            </Grid>
        </React.Fragment>
    );
}