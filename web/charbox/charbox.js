let charData
window.onload = winOnloadGenerate()

function getSelectClassNameValue() {
    if (document.getElementById("select_class_name").value === "Случайный класс") {
        return "random"
    }
    return document.getElementById("select_class_name").value;
}
function getSelectRaceNameValue() {
    if (document.getElementById("select_race_name").value === "Случайная раса") {
        return "random"
    }
    return document.getElementById("select_race_name").value;
}

async function winOnloadGenerate() {
    hideUploadPage()
    let req_json = `{"class":"random", "race":"random"}`

    const response = await fetch(`${window.location.protocol}//${window.location.hostname}:${window.location.port}/api/v1/post-current-character`,
        { method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(JSON.parse(req_json))
        });
    if (response.status === 429) {
        document.getElementById("lbl_429_warning").innerHTML = "Был превышен лимит вызова генерации!"
    }
    const json = await response.json();
    let data = JSON.stringify(json);

    charData = data
    WriteAllLabels(data)
}

async function getCharacter() {
    hideUploadPage()
    let req_json = `{"class":"${getSelectClassNameValue()}", "race":"${getSelectRaceNameValue()}"}`

    const response = await fetch(`${window.location.protocol}//${window.location.hostname}:${window.location.port}/api/v1/post-current-character`,
        { method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(JSON.parse(req_json))
        });
    if (response.status !== 200) {
        document.getElementById("lbl_429_warning").innerHTML = "Был превышен лимит вызова генерации!"
    }
    const json = await response.json();
    let data = JSON.stringify(json);

    charData = data
    WriteAllLabels(data)
}

async function WriteAllLabels(data) {
    await writeRacePhotoLabels(data)
    await writeToAbilitiesLabels(data)
    await writeToSaveThrowsLabels(data)
    await writeToSkillsLabels(data)
    await writeOtherLabels(data)
    await writeBackgroundLabels(data)
    await writeAppearanceLabels(data)
    await writeProficienciesLabels(data)
    await writeRaceAbilitiesLabels(data)
    await writeClassEquipmentLabels(data)
    await writeArmorLabels(data)
    await writeWeaponLabels(data)
}

function writeToAbilitiesLabels(data) {
    document.getElementById("lbl_className_info").innerHTML = JSON.parse(data)["class"]["class_name_ru"];

    let abilities = JSON.parse(data)["class"]["ability"];
    document.getElementById("lbl_ability_total_info").innerHTML = abilities["total"];
    document.getElementById("lbl_ability_strength_info").innerHTML = abilities["strength"];
    document.getElementById("lbl_ability_dexterity_info").innerHTML = abilities["dexterity"];
    document.getElementById("lbl_ability_bodyDifficulty_info").innerHTML = abilities["body_difficulty"];
    document.getElementById("lbl_ability_intelligence_info").innerHTML = abilities["intelligence"];
    document.getElementById("lbl_ability_wisdom_info").innerHTML = abilities["wisdom"];
    document.getElementById("lbl_ability_charisma_info").innerHTML = abilities["charisma"];

    let modifier = JSON.parse(data)["class"]["modifier"];
    document.getElementById("lbl_modifier_strength_info").innerHTML = modifier["strength"];
    document.getElementById("lbl_modifier_dexterity_info").innerHTML =  modifier["dexterity"];
    document.getElementById("lbl_modifier_bodyDifficulty_info").innerHTML = modifier["body_difficulty"];
    document.getElementById("lbl_modifier_intelligence_info").innerHTML =  modifier["intelligence"];
    document.getElementById("lbl_modifier_wisdom_info").innerHTML = modifier["wisdom"];
    document.getElementById("lbl_modifier_charisma_info").innerHTML = modifier["charisma"];
}

function writeToSaveThrowsLabels(data) {
    let saving_throws = JSON.parse(data)["class"]["saving_throws"];

    //обнуление радио
    var radios = ["rd_modifier_ST_strength_info", "rd_modifier_ST_dexterity_info",
        "rd_modifier_ST_bodyDifficulty_info", "rd_modifier_ST_intelligence_info",
        "rd_modifier_ST_wisdom_info", "rd_modifier_ST_charisma_info"];

    for (var i = 0; i < radios.length; i++) {
        document.getElementById(radios[i]).checked=false;
    }

    document.getElementById("lbl_modifier_ST_strength_info").innerHTML = saving_throws["strength"]["point"];
    document.getElementById("lbl_modifier_ST_dexterity_info").innerHTML =  saving_throws["dexterity"]["point"];
    document.getElementById("lbl_modifier_ST_bodyDifficulty_info").innerHTML = saving_throws["body_difficulty"]["point"];
    document.getElementById("lbl_modifier_ST_intelligence_info").innerHTML =  saving_throws["intelligence"]["point"];
    document.getElementById("lbl_modifier_ST_wisdom_info").innerHTML = saving_throws["wisdom"]["point"];
    document.getElementById("lbl_modifier_ST_charisma_info").innerHTML = saving_throws["charisma"]["point"];

    if (saving_throws["strength"]["mastery"] === true){
        document.getElementById("rd_modifier_ST_strength_info").checked=true;
    }
    if (saving_throws["dexterity"]["mastery"] === true){
        document.getElementById("rd_modifier_ST_dexterity_info").checked=true;
    }
    if (saving_throws["body_difficulty"]["mastery"]=== true){
        document.getElementById("rd_modifier_ST_bodyDifficulty_info").checked=true;
    }
    if (saving_throws["intelligence"]["mastery"]=== true){
        document.getElementById("rd_modifier_ST_intelligence_info").checked=true;
    }
    if (saving_throws["wisdom"]["mastery"]=== true){
        document.getElementById("rd_modifier_ST_wisdom_info").checked=true;
    }
    if (saving_throws["charisma"]["mastery"]=== true){
        document.getElementById("rd_modifier_ST_charisma_info").checked=true;
    }
}

function writeToSkillsLabels(data) {
    let skills =  JSON.parse(data)["skills"];

    //обнуление радио
    var radios = ["rd_Acrobatics", "rd_Animal_Handling","rd_Arcana","rd_Athletics",
        "rd_Deception","rd_History","rd_Insight","rd_Intimidation",
        "rd_Investigation","rd_Medicine","rd_Nature","rd_Perception",
        "rd_Performance","rd_Persuasion", "rd_Religion",
        "rd_SleightofHand", "rd_Stealth", "rd_Survival"];

    for (var i = 0; i < radios.length; i++) {
        document.getElementById(radios[i]).checked=false;
    }

    document.getElementById("lbl_Acrobatics").innerHTML = skills["acrobatics"]["modifier_value"];
    document.getElementById("lbl_Animal_Handling").innerHTML = skills["animal_handling"]["modifier_value"];
    document.getElementById("lbl_Arcana").innerHTML = skills["arcana"]["modifier_value"];
    document.getElementById("lbl_Athletics").innerHTML = skills["athletics"]["modifier_value"];
    document.getElementById("lbl_Deception").innerHTML = skills["deception"]["modifier_value"];
    document.getElementById("lbl_History").innerHTML = skills["history"]["modifier_value"];
    document.getElementById("lbl_Insight").innerHTML = skills["insight"]["modifier_value"];
    document.getElementById("lbl_Intimidation").innerHTML = skills["intimidation"]["modifier_value"];
    document.getElementById("lbl_Investigation").innerHTML = skills["investigation"]["modifier_value"];
    document.getElementById("lbl_Medicine").innerHTML = skills["medicine"]["modifier_value"];
    document.getElementById("lbl_Nature").innerHTML = skills["nature"]["modifier_value"]
    document.getElementById("lbl_Perception").innerHTML = skills["perception"]["modifier_value"];
    document.getElementById("lbl_Performance").innerHTML = skills["performance"]["modifier_value"];
    document.getElementById("lbl_Persuasion").innerHTML = skills["persuasion"]["modifier_value"];
    document.getElementById("lbl_Religion").innerHTML = skills["religion"]["modifier_value"];
    document.getElementById("lbl_SleightofHand").innerHTML = skills["sleight_of_hand"]["modifier_value"];
    document.getElementById("lbl_Stealth").innerHTML = skills["stealth"]["modifier_value"];
    document.getElementById("lbl_Survival").innerHTML = skills["survival"]["modifier_value"];


    if (skills["acrobatics"]["proficiency"] === true){
        document.getElementById("rd_Acrobatics").checked=true;
    }
    if (skills["animal_handling"]["proficiency"] === true){
        document.getElementById("rd_Animal_Handling").checked=true;
    }
    if (skills["arcana"]["proficiency"] === true){
        document.getElementById("rd_Arcana").checked=true;
    }
    if (skills["athletics"]["proficiency"] === true){
        document.getElementById("rd_Athletics").checked=true;
    }
    if (skills["deception"]["proficiency"] === true){
        document.getElementById("rd_Deception").checked=true;
    }
    if (skills["history"]["proficiency"] === true){
        document.getElementById("rd_History").checked=true;
    }
    if (skills["insight"]["proficiency"] === true){
        document.getElementById("rd_Insight").checked=true;
    }
    if (skills["intimidation"]["proficiency"] === true){
        document.getElementById("rd_Intimidation").checked=true;
    }
    if (skills["investigation"]["proficiency"] === true){
        document.getElementById("rd_Investigation").checked=true;
    }
    if (skills["medicine"]["proficiency"] === true){
        document.getElementById("rd_Medicine").checked=true;
    }
    if (skills["nature"]["proficiency"] === true){
        document.getElementById("rd_Nature").checked=true;
    }
    if (skills["perception"]["proficiency"] === true){
        document.getElementById("rd_Perception").checked=true;
    }
    if (skills["performance"]["proficiency"] === true){
        document.getElementById("rd_Performance").checked=true;
    }
    if (skills["persuasion"]["proficiency"] === true){
        document.getElementById("rd_Persuasion").checked=true;
    }
    if (skills["religion"]["proficiency"] === true){
        document.getElementById("rd_Religion").checked=true;
    }
    if (skills["sleight_of_hand"]["proficiency"] === true){
        document.getElementById("rd_SleightofHand").checked=true;
    }
    if (skills["stealth"]["proficiency"] === true){
        document.getElementById("rd_Stealth").checked=true;
    }
    if (skills["survival"]["proficiency"] === true){
        document.getElementById("rd_Survival").checked=true;
    }
}

function writeOtherLabels(data) {
    document.getElementById("lbl_passive_wisdom").innerHTML = JSON.parse(data)["passive_wisdom"];

    let clas = JSON.parse(data)["class"]
    document.getElementById("lbl_initiative").innerHTML = clas["initiative"];
    document.getElementById("lbl_hitsDice").innerHTML = clas["hits"]["hit_dice"];
    document.getElementById("lbl_hitsCount").innerHTML = clas["hits"]["hit_count"];
    document.getElementById("lbl_proficiency_bonus").innerHTML = clas["proficiency_bonus"];

    if (clas["inspiration"] === true) {
        document.getElementById("lbl_inspiration").innerHTML = "*" ;
    }
}

function writeBackgroundLabels(data) {
    let backgr = JSON.parse(data)["background"]
    document.getElementById("lbl_list_background_equipment").innerHTML = "";

    document.getElementById("lbl_background_name").innerHTML = backgr["background_name_ru"];
    document.getElementById("lbl_background_type").innerHTML = backgr["type"];
    document.getElementById("lbl_background_description").innerHTML = backgr["description"];
    document.getElementById("lbl_advice").innerHTML = backgr["advice"];
    document.getElementById("lbl_personalization").innerHTML = backgr["personalization"];
    document.getElementById("lbl_characterTrait").innerHTML = backgr["character_trait"];
    document.getElementById("lbl_ideal").innerHTML = backgr["ideal"];
    document.getElementById("lbl_worldview").innerHTML = backgr["worldview"];
    document.getElementById("lbl_affection").innerHTML = backgr["affection"];
    document.getElementById("lbl_weakness").innerHTML = backgr["weakness"];
    document.getElementById("lbl_gold").innerHTML = backgr["gold"];

    document.getElementById("lbl_background_ability").innerHTML = "<br> <strong>"+ backgr["background_ability"]["AbilityName"]+"</strong>"
        + " - " + backgr["background_ability"]["Description"];

    backgrEquip = backgr["equipment"]
    for(let i = 0; i < backgrEquip.length; i++){
        if (i !== backgrEquip.length-1) {
            document.getElementById("lbl_list_background_equipment").innerHTML += backgrEquip[i]+ ", " ;
        } else {
            document.getElementById("lbl_list_background_equipment").innerHTML += backgrEquip[i];
        }
    }

    console.log(document.getElementById("lbl_list_background_equipment").innerHTML)
    document.getElementById("lbl_background_instruments").innerHTML = backgr["mastery_of_tools"]

}

function writeAppearanceLabels(data) {
    let race = JSON.parse(data)["race"];
    document.getElementById("p_dragon_type").style.display = "none";
    document.getElementById("div_snake_appearance").style.visibility= "hidden";


    document.getElementById("lbl_raceName").innerHTML = race["race_type_name_ru"];
    document.getElementById("lbl_charFirstName").innerHTML =  race["first_name"];
    document.getElementById("lbl_charLastName").innerHTML =  race["last_name"];

    document.getElementById("lbl_charFirstName2").innerHTML =  race["first_name"];
    document.getElementById("lbl_charLastName2").innerHTML =  race["last_name"];
    document.getElementById("lbl_resist").innerHTML =  race["resist"];
    document.getElementById("lbl_bodySize").innerHTML =  race["body_size"];

    document.getElementById("lbl_speed").innerHTML =  race["speed"];
    document.getElementById("lbl_languages").innerHTML =  race["langs"];

    document.getElementById("lbl_gender").innerHTML =  race["gender"];
    document.getElementById("lbl_age").innerHTML =  race["age"];
    document.getElementById("lbl_eyesColor").innerHTML =  race["eyes"];
    document.getElementById("lbl_height").innerHTML =  race["height"];
    document.getElementById("lbl_weight").innerHTML =  race["weight"];

    document.getElementById("lbl_height_ft").innerHTML =  race["height_ft"];
    document.getElementById("lbl_weight_lb").innerHTML =  race["weight_lb"];

    document.getElementById("lbl_hairColor").innerHTML =  race["hair"];

    if (race["race_name"] === "Dragonborn"){
        document.getElementById("p_dragon_type").style.display = "inline";
        document.getElementById("lbl_dragon_color").innerHTML = race["other"]["dragon_type"]["color"];
    }

    if (race["race_name"] === "Yuan-ti"){
        document.getElementById("div_snake_appearance").style.visibility= "visible"
        document.getElementById("lbl_hairColor").innerHTML = "Нет";

        document.getElementById("lbl_typeSnakeBody").innerHTML =  race["other"]["yuanti_appearance"]["type_snake_body"];
        document.getElementById("lbl_humanoidSkinColor").innerHTML =  race["other"]["yuanti_appearance"]["humanoid_skin_color"];
        document.getElementById("lbl_scaleColor").innerHTML =  race["other"]["yuanti_appearance"]["scale_color"];
        document.getElementById("lbl_ScalePattern").innerHTML =  race["other"]["yuanti_appearance"]["scale_pattern"];
        document.getElementById("lbl_tongueColor").innerHTML =  race["other"]["yuanti_appearance"]["tongue_color"];
        document.getElementById("lbl_eyesColor").innerHTML =  race["other"]["yuanti_appearance"]["eye_color"];
        document.getElementById("lbl_headShape").innerHTML =  race["other"]["yuanti_appearance"]["head_shape"];
        document.getElementById("lbl_purebredsSpecialty").innerHTML =  race["other"]["yuanti_appearance"]["purebreds_specialty"];

    }
}

function writeProficienciesLabels(data) {
    let prof = JSON.parse(data)["class"]["proficiencies"];

    document.getElementById("lbl_weapon").innerHTML = prof["weapons"];
    document.getElementById("lbl_class_instruments").innerHTML = prof["instruments"];
    document.getElementById("lbl_armor").innerHTML = prof["armor"];
}

function writeRaceAbilitiesLabels(data) {
    document.getElementById("lbl_race_abilities").innerHTML = ""
        let race_abil = JSON.parse(data)["race"]["race_abilities"];

    if (race_abil === null){
        document.getElementById("lbl_race_abilities").innerHTML = "Нету"
        return
    }
    for(let i = 0; i < race_abil.length; i++){
        document.getElementById("lbl_race_abilities").innerHTML +=
            "<strong>"+ JSON.parse(JSON.stringify(race_abil[i]["AbilityName"]))+"</strong>" +": "
                + JSON.parse(JSON.stringify(race_abil[i]["Description"])) + "<br>" ;
    }
}

function writeClassEquipmentLabels(data) {
    document.getElementById("lbl_list_class_equipment").innerHTML = ""
    let equip = JSON.parse(data)["class"]["class_equipment"];

    for(let i = 0; i < equip.length; i++){
        if (i !== equip.length-1){
            comma = ", "
        }else {
            comma = ""
        }
        if (JSON.stringify(equip[i]["count"]) > 1){
            document.getElementById("lbl_list_class_equipment").innerHTML += JSON.parse(JSON.stringify(equip[i]["itemName"])) +
                " (" + JSON.parse(JSON.stringify(equip[i]["count"]))+ ")" + comma;
        } else {
            document.getElementById("lbl_list_class_equipment").innerHTML += JSON.parse(JSON.stringify(equip[i]["itemName"])) + comma;
        }
    }
}
function writeArmorLabels(data) {
    let armor = JSON.parse(data)["class"]["armor"];

    for(let i = 0; i < 1; i++) {
        document.getElementById("lbl_armor_name").innerHTML = JSON.parse(JSON.stringify(armor[i]["armorName"]));
        document.getElementById("lbl_armor_class").innerHTML = JSON.stringify(armor[i]["armorClassCount"]);
    }
}

function writeWeaponLabels(data) {
    document.getElementById("lbl_weapon_list").innerHTML = ""

    let weapon = JSON.parse(data)["class"]["weapon"];
    let count = " ";

    for(let i = 0; i < weapon.length; i++) {
        if (JSON.stringify(weapon[i]["count"]) > 1){
            count = " (" + JSON.stringify(weapon[i]["count"]) + ") "
        }
        document.getElementById("lbl_weapon_list").innerHTML += "[ "+JSON.stringify(weapon[i]["weaponName"]) + count
            + JSON.stringify(weapon[i]["damage"]) +" "+ JSON.stringify(weapon[i]["property"]) + " ]<br>";
    }
}

function writeRacePhotoLabels(data) {
    document.getElementById("img_Character_Preview").src = "";

    document.getElementById("img_Character_Preview").src = "imgs/"+ JSON.parse(data)["race"]["race_photo"]["path"] +
        JSON.parse(data)["race"]["race_photo"]["file_name"] ;
    document.getElementById("img_Character_Preview").alt = "Арт является примерным видом персонажа."
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

        WriteAllLabels(data)
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
    (async () => {
        const rawResponse = await fetch(`${window.location.protocol}//${window.location.hostname}:${window.location.port}/api/v1/run-lss-export`,
            { method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: charData
        });
        const content = await rawResponse.json();
        //console.log(JSON.stringify(content));

        var dataStr = "data:text/json;charset=utf-8," + encodeURIComponent(JSON.stringify(content));
        var downloadAnchorNode = document.createElement('a');
        downloadAnchorNode.setAttribute("href", dataStr);
        downloadAnchorNode.setAttribute("download",content["name"]["value"]+"_"+content["info"]["charClass"]["value"]+"_export_to_lss(beta).json");
        document.body.appendChild(downloadAnchorNode); // required for firefox
        downloadAnchorNode.click();
        downloadAnchorNode.remove();
    })();
}
