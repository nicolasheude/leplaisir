import React from 'react';
import Grid from '@material-ui/core/Grid';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import FormControl from '@material-ui/core/FormControl';
import FormLabel from '@material-ui/core/FormLabel';


export default function RadioButtonsGroupRepas() {
    const [value, setValue] = React.useState('oui');

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setValue((event.target as HTMLInputElement).value);
    };

    return (
        <Grid container spacing={3}>
            <FormControl component="fieldset"
            style={{
                margin: 25,
                minWidth: 120
                }}>
                <FormLabel component="legend">Repas Midi</FormLabel>
                <RadioGroup aria-label="answer" name="answer1" value={value} onChange={handleChange}>
                    <FormControlLabel value="restaurant" control={<Radio />} label="Restaurant Tennis" />
                    <FormControlLabel value="pic-nic" control={<Radio />} label="Pic-nic" />
                    <FormControlLabel value="autre" control={<Radio />} label="Autre" />
                </RadioGroup>
            </FormControl>
        </Grid>
    );
}

function RadioButtonsGroupRaquette() {
    const [value, setValue] = React.useState('oui');

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setValue((event.target as HTMLInputElement).value);
    };

    return (
        <Grid container spacing={3}>
            <FormControl component="fieldset"
            style={{
                margin: 25,
                minWidth: 120
                }}>
                <FormLabel component="legend">Location de raquette</FormLabel>
                <RadioGroup aria-label="answer" name="answer1" value={value} onChange={handleChange}>
                    <FormControlLabel value="oui" control={<Radio />} label="oui" />
                    <FormControlLabel value="non" control={<Radio />} label="non" />
                </RadioGroup>
            </FormControl>
        </Grid>
    );
}

function RadioButtonsGroupLicense() {
    const [value, setValue] = React.useState('oui');

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setValue((event.target as HTMLInputElement).value);
    };

    return (
        <Grid container spacing={3}>
            <FormControl component="fieldset"
            style={{
                margin: 25,
                minWidth: 120
                }}>
                <FormLabel component="legend">Licencié de la Fédération Française de Tennis (FFT)</FormLabel>
                <RadioGroup aria-label="answer" name="answer1" value={value} onChange={handleChange}>
                    <FormControlLabel value="oui" control={<Radio />} label="oui" />
                    <FormControlLabel value="non" control={<Radio />} label="non" />
                </RadioGroup>
            </FormControl>
        </Grid>
    );
}

function RadioButtonsGroupGender() {
    const [value, setValue] = React.useState('female');

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setValue((event.target as HTMLInputElement).value);
    };

    return (
        <Grid container spacing={3}>
            <FormControl component="fieldset"
            style={{
                margin: 25,
                minWidth: 120
                }}>
                <FormLabel component="legend">Sexe</FormLabel>
                <RadioGroup aria-label="gender" name="gender1" value={value} onChange={handleChange}>
                    <FormControlLabel value="feminin" control={<Radio />} label="Féminin" />
                    <FormControlLabel value="masculin" control={<Radio />} label="Masculin" />
                </RadioGroup>
            </FormControl>
        </Grid>
    );
}