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
            <FormControl component="fieldset"
            style={{
                margin: 25,
                minWidth: 120
                }}>
                <FormLabel component="legend">Je certifie être :</FormLabel>
                <RadioGroup aria-label="answer" name="answer1" value={value} onChange={handleChange}>
                    <FormControlLabel value="pere" control={<Radio />} label="Le père de l'enfant" />
                    <FormControlLabel value="mere" control={<Radio />} label="La mère de l'enfant" />
                    <FormControlLabel value="grandparent" control={<Radio />} label="Le grand parent de l'enfant" />
                </RadioGroup>

                <FormLabel component="legend">Autorise le club à prendre toutes les mesures d'urgence (hospitalisation).</FormLabel>
                <RadioGroup aria-label="answer" name="answer1" value={value} onChange={handleChange}>
                    <FormControlLabel value="oui" control={<Radio />} label="oui" />
                    <FormControlLabel value="non" control={<Radio />} label="non" />
                </RadioGroup>

                <FormLabel component="legend">Autorise l'équipe pédagogique à transporter mon enfant sur les lieux d'activités extra-tennis.</FormLabel>
                <RadioGroup aria-label="answer" name="answer1" value={value} onChange={handleChange}>
                    <FormControlLabel value="oui" control={<Radio />} label="oui" />
                    <FormControlLabel value="non" control={<Radio />} label="non" />
                </RadioGroup>

                <FormLabel component="legend">S'engage à faire vérifier par son médecin l'aptitude de l'enfant à la pratique du sport.</FormLabel>
                <RadioGroup aria-label="answer" name="answer1" value={value} onChange={handleChange}>
                    <FormControlLabel value="oui" control={<Radio />} label="oui" />
                    <FormControlLabel value="non" control={<Radio />} label="non" />
                </RadioGroup>

                <FormLabel component="legend">Autorise le Club à utiliser l'image de l'enfant sur son site internet et sur les documents de communication.</FormLabel>
                <RadioGroup aria-label="answer" name="answer1" value={value} onChange={handleChange}>
                    <FormControlLabel value="oui" control={<Radio />} label="oui" />
                    <FormControlLabel value="non" control={<Radio />} label="non" />
                </RadioGroup>


            </FormControl>
        </Grid>
    );
}

/*
S'engage à avoir pris connaissance des "conditions de stage"
*/