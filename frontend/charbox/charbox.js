
window.onload = getCurrentClass();
let charData

async function getCurrentClass() {
    hideUploadPage()
    const response = await fetch(`${window.location.protocol}//${window.location.hostname}:${window.location.port}/api/v1/get-character`);
    const json = await response.json();
    let data = JSON.stringify(json);

    charData = data

    //console.log(data)
    writeToAbilitiesLabels(data)
    writeToSaveThrowsLabels(data)
    writeToSkillsLabels(data)
    writeOtherLabels(data)
    writeBackgroundLabels(data)
    writeAppearanceLabels(data)
    writeProficiencies(data)
    writeRaceAbilities(data)
}

function writeToAbilitiesLabels(data) {
    document.getElementById("lbl_className_info").innerHTML = JSON.parse(data)["class"]["class_name_ru"];

    let abilities = JSON.stringify(JSON.parse(data)["class"]["ability"]);
    document.getElementById("lbl_ability_total_info").innerHTML = JSON.parse(abilities)["total"];
    document.getElementById("lbl_ability_strength_info").innerHTML = JSON.parse(abilities)["strength"];
    document.getElementById("lbl_ability_dexterity_info").innerHTML = JSON.parse(abilities)["dexterity"];
    document.getElementById("lbl_ability_bodyDifficulty_info").innerHTML = JSON.parse(abilities)["body_difficulty"];
    document.getElementById("lbl_ability_intelligence_info").innerHTML = JSON.parse(abilities)["intelligence"];
    document.getElementById("lbl_ability_wisdom_info").innerHTML = JSON.parse(abilities)["wisdom"];
    document.getElementById("lbl_ability_charisma_info").innerHTML = JSON.parse(abilities)["charisma"];

    let modifier = JSON.stringify(JSON.parse(data)["class"]["modifier"]);
    document.getElementById("lbl_modifier_strength_info").innerHTML = JSON.parse(modifier)["strength"];
    document.getElementById("lbl_modifier_dexterity_info").innerHTML =  JSON.parse(modifier)["dexterity"];
    document.getElementById("lbl_modifier_bodyDifficulty_info").innerHTML = JSON.parse(modifier)["body_difficulty"];
    document.getElementById("lbl_modifier_intelligence_info").innerHTML =  JSON.parse(modifier)["intelligence"];
    document.getElementById("lbl_modifier_wisdom_info").innerHTML = JSON.parse(modifier)["wisdom"];
    document.getElementById("lbl_modifier_charisma_info").innerHTML = JSON.parse(modifier)["charisma"];
}

function writeToSaveThrowsLabels(data) {
    let saving_throws = JSON.stringify(JSON.parse(data)["class"]["saving_throws"]);

    //обнуление радио
    var radios = ["rd_modifier_ST_strength_info", "rd_modifier_ST_dexterity_info",
        "rd_modifier_ST_bodyDifficulty_info", "rd_modifier_ST_intelligence_info",
        "rd_modifier_ST_wisdom_info", "rd_modifier_ST_charisma_info"];

    for (var i = 0; i < radios.length; i++) {
        document.getElementById(radios[i]).checked=false;
    }

    document.getElementById("lbl_modifier_ST_strength_info").innerHTML = JSON.parse(saving_throws)["strength"]["point"];
    document.getElementById("lbl_modifier_ST_dexterity_info").innerHTML =  JSON.parse(saving_throws)["dexterity"]["point"];
    document.getElementById("lbl_modifier_ST_bodyDifficulty_info").innerHTML = JSON.parse(saving_throws)["body_difficulty"]["point"];
    document.getElementById("lbl_modifier_ST_intelligence_info").innerHTML =  JSON.parse(saving_throws)["intelligence"]["point"];
    document.getElementById("lbl_modifier_ST_wisdom_info").innerHTML = JSON.parse(saving_throws)["wisdom"]["point"];
    document.getElementById("lbl_modifier_ST_charisma_info").innerHTML = JSON.parse(saving_throws)["charisma"]["point"];

    if (JSON.parse(saving_throws)["strength"]["mastery"] === true){
        document.getElementById("rd_modifier_ST_strength_info").checked=true;
    }
    if (JSON.parse(saving_throws)["dexterity"]["mastery"] === true){
        document.getElementById("rd_modifier_ST_dexterity_info").checked=true;
    }
    if (JSON.parse(saving_throws)["body_difficulty"]["mastery"]=== true){
        document.getElementById("rd_modifier_ST_bodyDifficulty_info").checked=true;
    }
    if (JSON.parse(saving_throws)["intelligence"]["mastery"]=== true){
        document.getElementById("rd_modifier_ST_intelligence_info").checked=true;
    }
    if (JSON.parse(saving_throws)["wisdom"]["mastery"]=== true){
        document.getElementById("rd_modifier_ST_wisdom_info").checked=true;
    }
    if (JSON.parse(saving_throws)["charisma"]["mastery"]=== true){
        document.getElementById("rd_modifier_ST_charisma_info").checked=true;
    }
}

function writeToSkillsLabels(data) {
    let skills = JSON.stringify(JSON.parse(data)["class"]["skills"]);

    //обнуление радио
    var radios = ["rd_Acrobatics", "rd_Animal_Handling","rd_Arcana","rd_Athletics",
        "rd_Deception","rd_History","rd_Insight","rd_Intimidation",
        "rd_Investigation","rd_Medicine","rd_Nature","rd_Perception",
        "rd_Performance","rd_Persuasion", "rd_Religion",
        "rd_SleightofHand", "rd_Stealth", "rd_Survival"];

    for (var i = 0; i < radios.length; i++) {
        document.getElementById(radios[i]).checked=false;
    }

    document.getElementById("lbl_Acrobatics").innerHTML = JSON.parse(skills)["acrobatics"]["modifier_value"];
    document.getElementById("lbl_Animal_Handling").innerHTML = JSON.parse(skills)["animal_handling"]["modifier_value"];
    document.getElementById("lbl_Arcana").innerHTML = JSON.parse(skills)["arcana"]["modifier_value"];
    document.getElementById("lbl_Athletics").innerHTML = JSON.parse(skills)["athletics"]["modifier_value"];
    document.getElementById("lbl_Deception").innerHTML = JSON.parse(skills)["deception"]["modifier_value"];
    document.getElementById("lbl_History").innerHTML = JSON.parse(skills)["history"]["modifier_value"];
    document.getElementById("lbl_Insight").innerHTML = JSON.parse(skills)["insight"]["modifier_value"];
    document.getElementById("lbl_Intimidation").innerHTML = JSON.parse(skills)["intimidation"]["modifier_value"];
    document.getElementById("lbl_Investigation").innerHTML = JSON.parse(skills)["investigation"]["modifier_value"];
    document.getElementById("lbl_Medicine").innerHTML = JSON.parse(skills)["medicine"]["modifier_value"];
    document.getElementById("lbl_Nature").innerHTML = JSON.parse(skills)["nature"]["modifier_value"]
    document.getElementById("lbl_Perception").innerHTML = JSON.parse(skills)["perception"]["modifier_value"];
    document.getElementById("lbl_Performance").innerHTML = JSON.parse(skills)["performance"]["modifier_value"];
    document.getElementById("lbl_Persuasion").innerHTML = JSON.parse(skills)["persuasion"]["modifier_value"];
    document.getElementById("lbl_Religion").innerHTML = JSON.parse(skills)["religion"]["modifier_value"];
    document.getElementById("lbl_SleightofHand").innerHTML = JSON.parse(skills)["sleight_of_hand"]["modifier_value"];
    document.getElementById("lbl_Stealth").innerHTML = JSON.parse(skills)["stealth"]["modifier_value"];
    document.getElementById("lbl_Survival").innerHTML = JSON.parse(skills)["survival"]["modifier_value"];


    if (JSON.parse(skills)["acrobatics"]["proficiency"] === true){
        document.getElementById("rd_Acrobatics").checked=true;
    }
    if (JSON.parse(skills)["animal_handling"]["proficiency"] === true){
        document.getElementById("rd_Animal_Handling").checked=true;
    }
    if (JSON.parse(skills)["arcana"]["proficiency"] === true){
        document.getElementById("rd_Arcana").checked=true;
    }
    if (JSON.parse(skills)["athletics"]["proficiency"] === true){
        document.getElementById("rd_Athletics").checked=true;
    }
    if (JSON.parse(skills)["deception"]["proficiency"] === true){
        document.getElementById("rd_Deception").checked=true;
    }
    if (JSON.parse(skills)["history"]["proficiency"] === true){
        document.getElementById("rd_History").checked=true;
    }
    if (JSON.parse(skills)["insight"]["proficiency"] === true){
        document.getElementById("rd_Insight").checked=true;
    }
    if (JSON.parse(skills)["intimidation"]["proficiency"] === true){
        document.getElementById("rd_Intimidation").checked=true;
    }
    if (JSON.parse(skills)["investigation"]["proficiency"] === true){
        document.getElementById("rd_Investigation").checked=true;
    }
    if (JSON.parse(skills)["medicine"]["proficiency"] === true){
        document.getElementById("rd_Medicine").checked=true;
    }
    if (JSON.parse(skills)["nature"]["proficiency"] === true){
        document.getElementById("rd_Nature").checked=true;
    }
    if (JSON.parse(skills)["perception"]["proficiency"] === true){
        document.getElementById("rd_Perception").checked=true;
    }
    if (JSON.parse(skills)["performance"]["proficiency"] === true){
        document.getElementById("rd_Performance").checked=true;
    }
    if (JSON.parse(skills)["persuasion"]["proficiency"] === true){
        document.getElementById("rd_Persuasion").checked=true;
    }
    if (JSON.parse(skills)["religion"]["proficiency"] === true){
        document.getElementById("rd_Religion").checked=true;
    }
    if (JSON.parse(skills)["sleight_of_hand"]["proficiency"] === true){
        document.getElementById("rd_SleightofHand").checked=true;
    }
    if (JSON.parse(skills)["stealth"]["proficiency"] === true){
        document.getElementById("rd_Stealth").checked=true;
    }
    if (JSON.parse(skills)["survival"]["proficiency"] === true){
        document.getElementById("rd_Survival").checked=true;
    }
}

function writeOtherLabels(data) {
    //document.getElementById("imageid").src="../template/save.png";


    document.getElementById("lbl_hitsDice").innerHTML = JSON.parse(data)["class"]["hits"]["hit_dice"];
    document.getElementById("lbl_hitsCount").innerHTML = JSON.parse(data)["class"]["hits"]["hit_count"];

    document.getElementById("lbl_proficiency_bonus").innerHTML = JSON.parse(data)["class"]["proficiency_bonus"];
    document.getElementById("lbl_passive_wisdom").innerHTML = JSON.parse(data)["class"]["passive_wisdom"];
    if (JSON.parse(data)["class"]["inspiration"] === true) {
        document.getElementById("lbl_inspiration").innerHTML = "*" ;
    }
}

function writeBackgroundLabels(data) {
    document.getElementById("lbl_background_name").innerHTML = JSON.parse(data)["background"]["background_name_ru"];
    document.getElementById("lbl_background_type").innerHTML = JSON.parse(data)["background"]["type"];
    document.getElementById("lbl_background_description").innerHTML = JSON.parse(data)["background"]["description"];
    document.getElementById("lbl_advice").innerHTML = JSON.parse(data)["background"]["advice"];
    document.getElementById("lbl_personalization").innerHTML = JSON.parse(data)["background"]["personalization"];
    document.getElementById("lbl_characterTrait").innerHTML = JSON.parse(data)["background"]["character_trait"];
    document.getElementById("lbl_ideal").innerHTML = JSON.parse(data)["background"]["ideal"];
    document.getElementById("lbl_worldview").innerHTML = JSON.parse(data)["background"]["worldview"];
    document.getElementById("lbl_affection").innerHTML = JSON.parse(data)["background"]["affection"];
    document.getElementById("lbl_weakness").innerHTML = JSON.parse(data)["background"]["weakness"];
}

function writeAppearanceLabels(data) {
    let race = JSON.stringify(JSON.parse(data)["race"]);
    document.getElementById("p_dragon_type").style.display = "none";
    document.getElementById("div_snake_appearance").style.display = "none";


    document.getElementById("lbl_raceName").innerHTML = JSON.parse(race)["race_type_name_ru"];
    document.getElementById("lbl_charFirstName").innerHTML =  JSON.parse(race)["first_name"];
    document.getElementById("lbl_charLastName").innerHTML =  JSON.parse(race)["last_name"];

    document.getElementById("lbl_charFirstName2").innerHTML =  JSON.parse(race)["first_name"];
    document.getElementById("lbl_charLastName2").innerHTML =  JSON.parse(race)["last_name"];
    document.getElementById("lbl_resist").innerHTML =  JSON.parse(race)["resist"];
    document.getElementById("lbl_bodySize").innerHTML =  JSON.parse(race)["body_size"];

    document.getElementById("lbl_speed").innerHTML =  JSON.parse(race)["speed"];
    document.getElementById("lbl_languages").innerHTML =  JSON.parse(race)["langs"];

    document.getElementById("lbl_gender").innerHTML =  JSON.parse(race)["gender"];
    document.getElementById("lbl_age").innerHTML =  JSON.parse(race)["age"];
    document.getElementById("lbl_eyesColor").innerHTML =  JSON.parse(race)["eyes"];
    document.getElementById("lbl_height").innerHTML =  JSON.parse(race)["height"];
    document.getElementById("lbl_weight").innerHTML =  JSON.parse(race)["weight"];

    document.getElementById("lbl_height_ft").innerHTML =  JSON.parse(race)["height_ft"];
    document.getElementById("lbl_weight_lb").innerHTML =  JSON.parse(race)["weight_lb"];

    document.getElementById("lbl_hairColor").innerHTML =  JSON.parse(race)["hair"];

    if (JSON.parse(race)["race_name"] === "Dragonborn"){
        document.getElementById("p_dragon_type").style.display = "inline";
        document.getElementById("lbl_dragon_color").innerHTML = JSON.parse(race)["other"]["dragon_type"]["color"];
    }

    if (JSON.parse(race)["race_name"] === "Yuan-ti"){
        document.getElementById("div_snake_appearance").style.display = "flex";
        document.getElementById("lbl_hairColor").innerHTML = "Нет";

        document.getElementById("lbl_typeSnakeBody").innerHTML =  JSON.parse(race)["other"]["yuanti_appearance"]["type_snake_body"];
        document.getElementById("lbl_humanoidSkinColor").innerHTML =  JSON.parse(race)["other"]["yuanti_appearance"]["humanoid_skin_color"];
        document.getElementById("lbl_scaleColor").innerHTML =  JSON.parse(race)["other"]["yuanti_appearance"]["scale_color"];
        document.getElementById("lbl_ScalePattern").innerHTML =  JSON.parse(race)["other"]["yuanti_appearance"]["scale_pattern"];
        document.getElementById("lbl_tongueColor").innerHTML =  JSON.parse(race)["other"]["yuanti_appearance"]["tongue_color"];
        document.getElementById("lbl_eyesColor").innerHTML =  JSON.parse(race)["other"]["yuanti_appearance"]["eye_color"];
        document.getElementById("lbl_headShape").innerHTML =  JSON.parse(race)["other"]["yuanti_appearance"]["head_shape"];
        document.getElementById("lbl_purebredsSpecialty").innerHTML =  JSON.parse(race)["other"]["yuanti_appearance"]["purebreds_specialty"];

    }
}

function writeProficiencies(data) {
    let prof = JSON.stringify(JSON.parse(data)["class"]["proficiencies"]);

    document.getElementById("lbl_weapon").innerHTML = JSON.parse(prof)["weapons"];
    document.getElementById("lbl_instruments").innerHTML = JSON.parse(prof)["instruments"];
    document.getElementById("lbl_armor").innerHTML = JSON.parse(prof)["armor"];
}

function writeRaceAbilities(data) {
    let race_abil = JSON.parse(data)["race"]["race_abilities"];

    for(let i = 0; i < race_abil.length; i++){
        console.log(race_abil[i])
        document.getElementById("lbl_race_abilities").innerHTML += "<strong>"+ JSON.stringify(race_abil[i]["AbilityName"])+"</strong>" +" --> "+ JSON.stringify(race_abil[i]["Description"]) + "<br>" ;
    }
}



const fileInput = document.querySelector('#file-js input[type=file]');
fileInput.onchange = () => {
    if (fileInput.files.length > 0) {
        const fileName = document.querySelector('#file-js .file-name');
        fileName.textContent = fileInput.files[0].name;
    }
}

function showUploadPage() {
    document.getElementById("file-js").style.display = "inline";
    document.getElementById("btn_show_upload").style.display = "none";
}

function hideUploadPage() {
    document.getElementById("file-js").style.display = "none";
    document.getElementById("btn_show_upload").style.display = "inline";
}

function importFile() {
    var files = document.getElementById('fl_upload_file').files;
    if (files.length <= 0) {
        return false;
    }
    var fr = new FileReader();
    fr.onload = function(e) {
        var result = JSON.parse(e.target.result);
        var data = JSON.stringify(result, null, 2);

        writeToAbilitiesLabels(data)
        writeToSaveThrowsLabels(data)
        writeToSkillsLabels(data)
        writeOtherLabels(data)
        writeBackgroundLabels(data)
        writeAppearanceLabels(data)
        writeProficiencies(data)
        writeRaceAbilities(data)
    };

    fr.readAsText(files.item(0));
}


function downloadObjectAsJson(){
    var dataStr = "data:text/json;charset=utf-8," + encodeURIComponent(JSON.stringify(JSON.parse(charData), null, 2));
    var downloadAnchorNode = document.createElement('a');
    downloadAnchorNode.setAttribute("href", dataStr);
    downloadAnchorNode.setAttribute("download","dnd_character.json");
    document.body.appendChild(downloadAnchorNode); // required for firefox
    downloadAnchorNode.click();
    downloadAnchorNode.remove();
}

function exportToLSS() {
    document.getElementById("btn_export_lss").innerHTML = "{Отсутствует интеграция c LSS}"
    document.getElementById("btn_export_lss").className = "button is-danger";
}