import * as React from 'react';

/*

    parent_nom: "",
    parent_prenom: "",
    parent_sexe: "",
    parent_adresse: "",
    parent_ville: "",
    parent_cp: "",
    parent_email: "",
    parent_tel_1: "",
    parent_tel_2: ""
*/
export default function Form() {
  const [state, setState] = React.useState({
    child_email: "",
    child_prenom: "",
    child_nom: "",
    child_anniversaire: "",
    child_nationalite: "",
    child_sexe: "",
    child_niveau: "",
    child_fft: "",
    child_locationR: "",
    child_repas: "",
    child_formule: "",
    child_semaine: "",

    parent_nom: "",
    parent_prenom: "",
    parent_sexe: "",
    parent_adresse: "",
    parent_ville: "",
    parent_cp: "",
    parent_email: "",
    parent_tel_1: "",
    parent_tel_2: ""
  })

  function handleChange(evt) {
    const value = evt.target.value;
    setState({
      ...state,
      [evt.target.name]: value
    });
  }  

  const handleSubmit = (event) => {
      let child_s = {
        email: event.target.child_email.value,
        nom: event.target.child_nom.value,
        prenom: event.target.child_prenom.value,
        anniversaire: event.target.child_anniversaire.value,
        nationalite: event.target.child_nationalite.value,
        sexe: event.target.child_sexe.value,
        niveau: event.target.child_niveau.value,
        fft: event.target.child_fft.value,
        locationR: event.target.child_locationR.value,
        repas: event.target.child_repas.value,
        formule: event.target.child_formule.value,
        semaine: event.target.child_semaine.value
      }

      let parents_s = {
          nom: event.target.parent_nom.value,
          prenom: event.target.parent_prenom.value,
          sexe: event.target.parent_sexe.value,
          adresse: event.target.parent_adresse.value,
          ville: event.target.parent_ville.value,
          cp: event.target.parent_cp.value,
          email: event.target.parent_email.value,
          tel_1: event.target.parent_tel_1.value,
          tel_2: event.target.parent_tel_2.value,
      }

      let body = {
          child: child_s,
          parents: parents_s
      }

      fetch('http://127.0.0.1:8080/hello', {
          method: 'POST',
          body: body
        }).then(function(response) {
//          console.log(response)
          return response;
        });
      event.preventDefault();
  }

  return (
    <React.Fragment>

      {/* Partie Gosse */}

      <h2>Votre enfant :</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Son email
          <input type="text" name="child_email" value={state.child_email} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Son nom
          <input type="text" name="child_nom" value={state.child_nom} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Son prénom
          <input type="text" name="child_prenom" value={state.child_prenom} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Sa date d'anniversaire
          <input type="text" name="child_anniversaire" value={state.child_anniversaire} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Sa nationalité
          <input type="text" name="child_nationalite" value={state.child_nationalite} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Son sexe
          <input type="text" name="child_sexe" value={state.child_sexe} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Son niveau
          <input type="text" name="child_niveau" value={state.child_niveau} onChange={handleChange} />
        </label>
        <br/>
        <label>
          FFT
          <input type="text" name="child_fft" value={state.child_fft} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Location Raquette ?
          <input type="text" name="child_locationR" value={state.child_locationR} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Repas ?
          <input type="text" name="child_repas" value={state.child_repas} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Formule (Atelier choisi)
          <input type="text" name="child_formule" value={state.child_formule} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Semaine
          <input type="text" name="child_semaine" value={state.child_semaine} onChange={handleChange} />
        </label>
        <br/>

        {/* Partie Parents */}

        <h2>Vous :</h2>

        <label>
          Votre nom
          <input type="text" name="parent_nom" value={state.parent_nom} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Votre prénom
          <input type="text" name="parent_prenom" value={state.parent_prenom} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Votre sexe
          <input type="text" name="parent_sexe" value={state.parent_sexe} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Votre adresse
          <input type="text" name="parent_adresse" value={state.parent_adresse} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Votre ville
          <input type="text" name="parent_ville" value={state.parent_ville} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Votre Code Postal
          <input type="text" name="parent_cp" value={state.parent_cp} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Votre Email
          <input type="text" name="parent_email" value={state.parent_email} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Votre Téléphone mobile
          <input type="text" name="parent_tel_1" value={state.parent_tel_1} onChange={handleChange} />
        </label>
        <br/>
        <label>
          Votre Téléphone fixe
          <input type="text" name="parent_tel_2" value={state.parent_tel_2} onChange={handleChange} />
        </label>
        <br/>

        <input type="submit" value="Envoyer" />
      </form>
    </React.Fragment>
  );
}
