import React from 'react';
import Grid from '@material-ui/core/Grid';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import FormControl from '@material-ui/core/FormControl';
import FormLabel from '@material-ui/core/FormLabel';

export default function RadioButtonsGroupRaquette() {
    const [value, setValue] = React.useState('oui');

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setValue((event.target as HTMLInputElement).value);
    };

    return (
        <Grid container spacing={3}>
                <FormLabel component="legend">Location de raquette</FormLabel>
                <RadioGroup aria-label="answer" name="answer1" value={value} onChange={handleChange}>
                    <FormControlLabel value="oui" control={<Radio />} label="oui" />
                    <FormControlLabel value="non" control={<Radio />} label="non" />
                </RadioGroup>
        </Grid>
    );
}
