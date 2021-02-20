import React from 'react';
import Grid from '@material-ui/core/Grid';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import FormControl from '@material-ui/core/FormControl';
import FormLabel from '@material-ui/core/FormLabel';

export default function RadioButtonsGroupGender() {
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