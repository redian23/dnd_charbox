let charData
document.addEventListener('DOMContentLoaded', (event) => {
    getCharacter();
});

function getSelectClassNameValue() {
    if (document.getElementById("select_class_name").value === "Случайный класс") {
        return "random";
    }
    return document.getElementById("select_class_name").value;
}

function getSelectClassArchetypeNameValue() {
    return document.getElementById("select_class_archetype").value;
}

////////////////////////////////////////////////////////////////////
function getSelectRaceNameValue() {
    if (document.getElementById("select_race_name").value === "Случайная раса") {
        return "random"
    }
    return document.getElementById("select_race_name").value;
}

function getSelectRaceArchetypeNameValue() {
    return document.getElementById("select_race_archetype").value;
}

///////////////////////////////////////////////////////////////////
function getSelectBackgroundNameValue() {
    if (document.getElementById("select_background_name").value === "Случайная предыстория") {
        return "random"
    }
    return document.getElementById("select_background_name").value;
}

function getSelectLevelValue() {
    if (document.getElementById("select_char_level").value === "1 уровень") {
        return 1
    }
    return document.getElementById("select_char_level").value;
}

function getSelectGenderValue() {
    return document.getElementById("select_gender").value;
}

document.addEventListener('DOMContentLoaded', applySavedColor);

async function getCharacter() {
    let req_json = {
        "class": getSelectClassNameValue(),
        "class_archetype": getSelectClassArchetypeNameValue(),
        "race": getSelectRaceNameValue(),
        "race_type": getSelectRaceArchetypeNameValue(),
        "background": getSelectBackgroundNameValue(),
        "level": parseInt(getSelectLevelValue(), 10),
        "gender": getSelectGenderValue()
    };
     console.log(JSON.parse(JSON.stringify(req_json)))

    const response = await fetch(`${window.location.protocol}//${window.location.hostname}:${window.location.port}/api/v2/get-character`,
        {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(req_json)
        });
    const json = await response.json();
    let data = JSON.stringify(json);

    charData = data;
    WriteAllLabels(data);

    console.log(json);
}

//////////////////////////////////////////////////////////////////////////////////
async function WriteAllLabels(data) {
    await writeToCharacterCardInputs(data) //done
    await writeToAbilitiesLabels(data) //done
    await writeToSaveThrowsLabels(data) //done
    await writeToSkillsLabels(data) //done

    await writeDebugSkillLabels(data) //TODO

    await writeToHitsLabels(data) //done
    await writeToPassiveInfoLabels(data) //done
    await writeOtherLabels(data) //done
    await writeProficienciesLabels(data)//done
    await writeClassAbilitiesLabels(data)//done
    await writeRaceAbilitiesLabels(data)//done
    await writeBackgroundAbilitiesLabels(data) //done
    await writeClassEquipmentLabels(data) //done
    await writeBackgroundEquipmentLabels(data) //done
    await writeWeaponLabels(data)
    await writeBackgroundLabels(data)
    await writeAppearanceLabels(data)
    await writeSpellCastingLabels(data)
    await writeSpellsLabels(data)
}

function writeToCharacterCardInputs(data) {
    document.getElementById("lbl_character_full_name").innerHTML = JSON.parse(data)["race"]["first_name"] + " " + JSON.parse(data)["race"]["last_name"] ;
    document.getElementById("lbl_level_info").innerHTML = JSON.parse(data)["level"]
    document.getElementById("lbl_exp_info").innerHTML = JSON.parse(data)["experience"]
    document.getElementById("lbl_class_name").innerHTML = JSON.parse(data)["class"]["class_name_ru"];
    document.getElementById("lbl_archetype_class_name").innerHTML = JSON.parse(data)["class"]["class_archetype_name"];
    document.getElementById("lbl_race_name").innerHTML = JSON.parse(data)["race"]["race_name_ru"];
    document.getElementById("lbl_background_name").innerHTML = JSON.parse(data)["background"]["background_name_ru"];
    document.getElementById("lbl_background_type").innerHTML = JSON.parse(data)["background"]["background_specific_type"];
    document.getElementById("lbl_worldview").innerHTML = JSON.parse(data)["background"]["personalization"]["worldview"];
}

function writeToAbilitiesLabels(data) {

    let abilities = JSON.parse(data)["class"]["ability"];
    document.getElementById("lbl_ability_strength_info").innerHTML = abilities["strength"];
    document.getElementById("lbl_ability_dexterity_info").innerHTML = abilities["dexterity"];
    document.getElementById("lbl_ability_body_difficulty_info").innerHTML = abilities["body_difficulty"];
    document.getElementById("lbl_ability_intelligence_info").innerHTML = abilities["intelligence"];
    document.getElementById("lbl_ability_wisdom_info").innerHTML = abilities["wisdom"];
    document.getElementById("lbl_ability_charisma_info").innerHTML = abilities["charisma"];

    let modifier = JSON.parse(data)["class"]["modifier"];
    document.getElementById("lbl_modifier_strength_info").innerHTML = modifier["strength"] > 0 ? `+${modifier["strength"]}` : modifier["strength"];
    document.getElementById("lbl_modifier_dexterity_info").innerHTML = modifier["dexterity"] > 0 ? `+${modifier["dexterity"]}` : modifier["dexterity"];
    document.getElementById("lbl_modifier_body_difficulty_info").innerHTML = modifier["body_difficulty"] > 0 ? `+${modifier["body_difficulty"]}` : modifier["body_difficulty"];
    document.getElementById("lbl_modifier_intelligence_info").innerHTML = modifier["intelligence"] > 0 ? `+${modifier["intelligence"]}` : modifier["intelligence"];
    document.getElementById("lbl_modifier_wisdom_info").innerHTML = modifier["wisdom"] > 0 ? `+${modifier["wisdom"]}` : modifier["wisdom"];
    document.getElementById("lbl_modifier_charisma_info").innerHTML = modifier["charisma"] > 0 ? `+${modifier["charisma"]}` : modifier["charisma"];
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
    //ПЕРЕДЕЛАТЬ *Звуки непереводимой смеси мата и айтишных терминов*
    let skills =  JSON.parse(data)["skills"];

    //обнуление радио
    const radios = ["rd_Acrobatics", "rd_Animal_Handling","rd_Arcana","rd_Athletics",
        "rd_Deception","rd_History","rd_Insight","rd_Intimidation",
        "rd_Investigation","rd_Medicine","rd_Nature","rd_Perception",
        "rd_Performance","rd_Persuasion", "rd_Religion",
        "rd_SleightofHand", "rd_Stealth", "rd_Survival"];

    for (let i = 0; i < radios.length; i++) {
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

    const dbl_radios = ["rd_Acrobatics_dbl_prof", "rd_Animal_Handling_dbl_prof","rd_Arcana_dbl_prof","rd_Athletics_dbl_prof",
        "rd_Deception_dbl_prof","rd_History_dbl_prof","rd_Insight_dbl_prof","rd_Intimidation_dbl_prof",
        "rd_Investigation_dbl_prof","rd_Medicine_dbl_prof","rd_Nature_dbl_prof","rd_Perception_dbl_prof",
        "rd_Performance_dbl_prof","rd_Persuasion_dbl_prof", "rd_Religion_dbl_prof",
        "rd_SleightofHand_dbl_prof", "rd_Stealth_dbl_prof", "rd_Survival_dbl_prof"];

    for (let i = 0; i < dbl_radios.length; i++) {
        document.getElementById(dbl_radios[i]).checked=false;
    }

    if (skills["acrobatics"]["double_proficiency"] === true){
        document.getElementById("rd_Acrobatics_dbl_prof").checked=true;
    }
    if (skills["animal_handling"]["double_proficiency"] === true){
        document.getElementById("rd_Animal_Handling_dbl_prof").checked=true;
    }
    if (skills["arcana"]["double_proficiency"] === true){
        document.getElementById("rd_Arcana_dbl_prof").checked=true;
    }
    if (skills["athletics"]["double_proficiency"] === true){
        document.getElementById("rd_Athletics_dbl_prof").checked=true;
    }
    if (skills["deception"]["double_proficiency"] === true){
        document.getElementById("rd_Deception_dbl_prof").checked=true;
    }
    if (skills["history"]["double_proficiency"] === true){
        document.getElementById("rd_History_dbl_prof").checked=true;
    }
    if (skills["insight"]["double_proficiency"] === true){
        document.getElementById("rd_Insight_dbl_prof").checked=true;
    }
    if (skills["intimidation"]["double_proficiency"] === true){
        document.getElementById("rd_Intimidation_dbl_prof").checked=true;
    }
    if (skills["investigation"]["double_proficiency"] === true){
        document.getElementById("rd_Investigation_dbl_prof").checked=true;
    }
    if (skills["medicine"]["double_proficiency"] === true){
        document.getElementById("rd_Medicine_dbl_prof").checked=true;
    }
    if (skills["nature"]["double_proficiency"] === true){
        document.getElementById("rd_Nature_dbl_prof").checked=true;
    }
    if (skills["perception"]["double_proficiency"] === true){
        document.getElementById("rd_Perception_dbl_prof").checked=true;
    }
    if (skills["performance"]["double_proficiency"] === true){
        document.getElementById("rd_Performance_dbl_prof").checked=true;
    }
    if (skills["persuasion"]["double_proficiency"] === true){
        document.getElementById("rd_Persuasion_dbl_prof").checked=true;
    }
    if (skills["religion"]["double_proficiency"] === true){
        document.getElementById("rd_Religion_dbl_prof").checked=true;
    }
    if (skills["sleight_of_hand"]["double_proficiency"] === true){
        document.getElementById("rd_SleightofHand_dbl_prof").checked=true;
    }
    if (skills["stealth"]["double_proficiency"] === true){
        document.getElementById("rd_Stealth_dbl_prof").checked=true;
    }
    if (skills["survival"]["double_proficiency"] === true){
        document.getElementById("rd_Survival_dbl_prof").checked=true;
    }
}

function writeDebugSkillLabels(data) {
    document.getElementById("lbl_race_skills").innerHTML = JSON.parse(data)["race"]["race_skill"].join(', ');
    document.getElementById("lbl_background_skills").innerHTML = JSON.parse(data)["background"]["background_skills"].join(', ');
    document.getElementById("lbl_class_skills").innerHTML = JSON.parse(data)["class"]["proficiencies"]["skills_of_class"].join(', ');
}

function writeToHitsLabels(data) {
    document.getElementById("lbl_hit_dice_count").innerHTML =  JSON.parse(data)["level"];
    document.getElementById("lbl_hitsDice").innerHTML = JSON.parse(data)["class"]["hits"]["hit_dice"];
    document.getElementById("lbl_hitsCount").innerHTML = JSON.parse(data)["class"]["hits"]["hit_count"];
}

function writeToPassiveInfoLabels(data) {
    document.getElementById("lbl_passive_wisdom_insight").innerHTML = JSON.parse(data)["passive_info"]["passive_wisdom_insight"];
    document.getElementById("lbl_passive_wisdom_perception").innerHTML = JSON.parse(data)["passive_info"]["passive_wisdom_perception"];
    document.getElementById("passive_intellect_investigation").innerHTML = JSON.parse(data)["passive_info"]["passive_intellect_investigation"];
}

function writeOtherLabels(data) {
    document.getElementById("lbl_proficiency_bonus").innerHTML = JSON.parse(data)["class"]["proficiency_bonus"];

    if (JSON.parse(data)["class"]["inspiration"] === true) {
        document.getElementById("lbl_inspiration").innerHTML = "1" ;
    }else {
        document.getElementById("lbl_inspiration").innerHTML = "Нет" ;
    }
    document.getElementById("lbl_armor_name").innerHTML = JSON.parse(data)["class"]["armor_info"]["armorName"]
    document.getElementById("lbl_armor_class").innerHTML = JSON.parse(data)["class"]["armor_info"]["armorClassCount"]
    document.getElementById("lbl_speed").innerHTML = JSON.parse(data)["race"]["race_type_name"]["speed"];
}

function writeProficienciesLabels(data) {
    document.getElementById("lbl_languages").innerHTML = JSON.parse(data)["langs"].join(', ');
    document.getElementById("lbl_languages").innerHTML = JSON.parse(data)["langs"].join(', ');
    document.getElementById("lbl_resist").innerHTML =  JSON.parse(data)["race"]["resist"].join(', ');
    document.getElementById("lbl_armor").innerHTML = JSON.parse(data)["class"]["proficiencies"]["armor"].join("\,\r\n");
    document.getElementById("lbl_weapon").innerHTML = JSON.parse(data)["class"]["proficiencies"]["weapons"].join(', ');
    document.getElementById("lbl_class_instruments").innerHTML = JSON.parse(data)["class"]["proficiencies"]["tools"].join(', ');
    //document.getElementById("lbl_class_skills").innerHTML = JSON.parse(data)["class"]["proficiencies"]["skills_of_class"].join(', ');
    document.getElementById("lbl_background_instruments").innerHTML = JSON.parse(data)["background"]["mastery_of_tools"].join(', ');
}

function writeClassAbilitiesLabels(data) {
    document.getElementById("lbl_class_abilities").innerHTML = ""
    let class_abil = JSON.parse(data)["class"]["class_abilities"];

    for(let i = 0; i < class_abil.length; i++){
        document.getElementById("lbl_class_abilities").innerHTML +=
            "<article><strong>"+ JSON.parse(JSON.stringify(class_abil[i]["name"]))+"</strong>" +": "
            + JSON.parse(JSON.stringify(class_abil[i]["description"])) + "</article>" ;
    }
}

function writeRaceAbilitiesLabels(data) {
    document.getElementById("lbl_race_abilities").innerHTML = "";

    race_abilities_json = JSON.parse(data)["race"]["race_type_name"]["race_abilities"]
    for (var ability in race_abilities_json) {
        document.getElementById("lbl_race_abilities").innerHTML +=
            "<article><strong>"+ ability +"</strong>" +": "+ race_abilities_json[ability]+ "</article>" ;
    }
}

function writeBackgroundAbilitiesLabels(data) {
    document.getElementById("lbl_background_ability").innerHTML = "";

    background_abilities_json = JSON.parse(data)["background"]["background_ability"]

    for (var ability in background_abilities_json) {
        document.getElementById("lbl_background_ability").innerHTML +=
            "<article><strong>"+ ability +"</strong>" +": "+ background_abilities_json[ability] + "</article>" ;
    }
}

function writeClassEquipmentLabels(data) {
    let classEquip = JSON.parse(data)["class"]["equipment"];

    var tableBody = document.getElementById("class_equipment_table_body");
    var row = "<tr>";
    for (let i in classEquip) {
        row += "<td>" + classEquip[i].name + "</td>";
        row += "<td>" + classEquip[i].count + "</td>";
        row += "</tr>";
    }
    tableBody.innerHTML = row;
}

function writeBackgroundEquipmentLabels(data) {
    backgrEquip = JSON.parse(data)["background"]["equipment"]

    var tableBody = document.getElementById("background_equipment_table_body");
    var row = "<tr>";
    for (let i in backgrEquip) {
        row += "<td>" + backgrEquip[i].name + "</td>";
        row += "<td>" + backgrEquip[i].count + "</td>";
        row += "</tr>";
    }
    tableBody.innerHTML = row;
}

function writeWeaponLabels(data) {
    weaponData = JSON.parse(data)["class"]["weapon_info"];

    var tableBody = document.getElementById("weapon_table_body");
    var row = "<tr>";
    for (let i in weaponData) {
        row += "<td>" + weaponData[i].weaponName + "</td>";
        row += "<td>" + weaponData[i].weaponType + "</td>";
        row += "<td>" + weaponData[i].damage + "</td>";
        row += "<td>" + weaponData[i].property + "</td>";
        row += "</tr>";
    }
    tableBody.innerHTML = row;
}

function writeBackgroundLabels(data) {
    let backgr = JSON.parse(data)["background"]

    document.getElementById("lbl_background_description").innerHTML = backgr["description"];
    document.getElementById("lbl_advice").innerHTML = backgr["personalization"]["advice"];
    document.getElementById("lbl_personalization").innerHTML = backgr["personalization"]["personalization_description"];
    document.getElementById("lbl_characterTrait").innerHTML = backgr["personalization"]["character_trait"];
    document.getElementById("lbl_ideal").innerHTML = backgr["personalization"]["ideal"];
    document.getElementById("lbl_affection").innerHTML = backgr["personalization"]["affection"];
    document.getElementById("lbl_weakness").innerHTML = backgr["personalization"]["weakness"];
}

function writeAppearanceLabels(data) {
    let race = JSON.parse(data)["race"];

    document.getElementById("lbl_raceName").innerHTML = race["race_name_ru"];
    document.getElementById("lbl_charFirstName").innerHTML =  race["first_name"];
    document.getElementById("lbl_charLastName").innerHTML =  race["last_name"];
    document.getElementById("lbl_gender").innerHTML =  race["gender"];

    document.getElementById("lbl_age").innerHTML =  race["body"]["age"];
    document.getElementById("lbl_bodySize").innerHTML =  race["body"]["body_size"];
    document.getElementById("lbl_eyesColor").innerHTML =  race["body"]["eyes"];
    document.getElementById("lbl_hairColor").innerHTML =  race["body"]["hair"];

    document.getElementById("lbl_height").innerHTML =  race["body"]["height"];
    document.getElementById("lbl_weight").innerHTML =  race["body"]["weight"];
    document.getElementById("lbl_height_ft").innerHTML =  ["body"]["height_ft"];
    document.getElementById("lbl_weight_lb").innerHTML =  race["body"]["weight_lb"];
}

function writeSpellCastingLabels(data) {
    let spell_using = JSON.parse(data)["class"]["spell_casting"];

    document.getElementById("lbl_basic_spell_characteristics").innerHTML =spell_using["basic_spell_characteristics"]
    document.getElementById("lbl_saving_throw_difficulty").innerHTML = spell_using["saving_throw_difficulty"]
    document.getElementById("lbl_spell_damage_modifier").innerHTML = spell_using["spell_damage_modifier"]
}

function writeSpellsLabels(data) {
    let spells_list = JSON.parse(data)["spells"];
    let spellLevels = [0, 1, 2, 3, 4];
    let spellLevelTables = {
        0: "zero_lvl_spell_table_body",
        1: "one_lvl_spell_table_body",
        2: "two_lvl_spell_table_body",
        3: "three_lvl_spell_table_body",
        4: "four_lvl_spell_table_body"
    };

    for (let level of spellLevels) {
        let spellsOfLevel = spells_list.filter(spell => spell.spellLevel === level);
        let tableBody = document.getElementById(spellLevelTables[level]);

        // Очистка таблицы перед добавлением новых строк
        tableBody.innerHTML = "";

        if (spellsOfLevel.length === 0) {
            tableBody.innerHTML = "<tr><td>Нет заклинаний данного уровня</td></tr>";
        } else {
            for (let spell of spellsOfLevel) {
                let row = "<tr>";
                row += "<td><a href="+spell.url+">" + spell.spellNameRu + " [" + spell.spellNameEng + "]"+ "</a></td>";
                row += "</tr>";
                tableBody.innerHTML += row;
            }
        }
    }
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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

function gotoFAQ(){
    window.location = `${window.location.protocol}//${window.location.hostname}:${window.location.port}/faq`;
}

document.addEventListener("DOMContentLoaded", function() {
    document.getElementById("select_class_name").addEventListener("change", updateClassArchetypeSelect);
    // Initial update to populate archetypes on page load if necessary
    updateClassArchetypeSelect();
});

document.addEventListener("DOMContentLoaded", function() {
    document.getElementById("select_char_level").addEventListener("change", updateClassArchetypeSelect);
    // Initial update to populate archetypes on page load if necessary
    updateClassArchetypeSelect();
});

document.addEventListener("DOMContentLoaded", function() {
    document.getElementById("select_race_name").addEventListener("change", updateRaceArchetypeSelect);
    // Initial update to populate archetypes on page load if necessary
    updateRaceArchetypeSelect();
});

document.addEventListener('DOMContentLoaded', function() {
    const themeSwitch = document.getElementById('theme-switch');

    // Устанавливаем тему при загрузке страницы
    setInitialTheme();

    themeSwitch.addEventListener('change', function() {
        const theme = themeSwitch.checked ? 'light' :'dark';
        setTheme(theme);
    });

    function setInitialTheme() {
        // Проверяем, какая тема выбрана (если выбрана) при загрузке страницы
        const savedTheme = localStorage.getItem('theme');
        if (savedTheme) {
            document.documentElement.setAttribute('data-theme', savedTheme);
            // Проверяем переключатель (checkbox) в соответствии с выбранной темой
            themeSwitch.checked = savedTheme === 'light';
        }
    }

    function setTheme(theme) {
        document.documentElement.setAttribute('data-theme', theme);
        // Сохраняем выбранную тему в localStorage
        localStorage.setItem('theme', theme);
    }
});


function changeColor(color) {
    const link = document.querySelector('link[rel*="stylesheet"]');
    link.href = `https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.${color}.min.css`;
    localStorage.setItem('color', color);
}

function applySavedColor() {
    const savedColor = localStorage.getItem('color');
    if (savedColor) {
        const link = document.querySelector('link[rel*="stylesheet"]');
        link.href = `https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.${savedColor}.min.css`;
    }
}

function addOptions(select, optionsArray) {
    optionsArray.forEach(optionText => {
        const option = document.createElement('option');
        option.value = optionText;
        option.textContent = optionText;
        select.appendChild(option);
    });
}

function updateClassArchetypeSelect() {
    let className = getSelectClassNameValue();
    let level = getSelectLevelValue()

    let classArchetype = document.getElementById('select_class_archetype');
    classArchetype.innerHTML = '';

    if (level < 3 || className === 'random'){
        classArchetype.style.display = "none";
    }else {
        classArchetype.style.display = "block";
    }

    switch (className) {
        case "Бард":
            addOptions(classArchetype, [ 'Случайная коллегия',
                'Коллегия доблести', 'Коллегия знаний', 'Коллегия мечей', 'Коллегия очарования',
                'Коллегия шёпотов', 'Коллегия красноречия', 'Коллегия созидания', 'Коллегия духов'
            ]);
            break;
        case "Варвар":
            addOptions(classArchetype, ['Случайный путь',
                'Путь берсерка', 'Путь тотемного воина', 'Путь буревестника', 'Путь предка-хранителя',
                'Путь фанатика', 'Путь дикой магии', 'Путь зверя', 'Путь великана', 'Путь бушующего в бою'
            ]);
            break;
        default:
            break;
    }
}


function updateRaceArchetypeSelect() {
    let raceName = getSelectRaceNameValue();

    let raceArchetype = document.getElementById('select_race_archetype');
    raceArchetype.innerHTML = '';

    if (raceName === 'random'){
        raceArchetype.style.display = "none";
    }else {
        raceArchetype.style.display = "block";
    }

    switch (raceName) {
        case "random":
            raceArchetype.style.display = "none";
            break;
        case "Аасимар":
            addOptions(raceArchetype, ["Аасимар-защитник","Аасимар–каратель","Аасимар–падший"]);
            break;
        case "Багбир":
            addOptions(raceArchetype, ["Стандартный"]);
            break;
        case "Гном":
            addOptions(raceArchetype, ["Лесной гном","Скальный гном"]);
            break;
        case "Гоблин":
            addOptions(raceArchetype, ["Стандартный","Гоблин (RLW)","Гоблин (GGR)","Гоблин Данквуда"]);
            break;
        case "Голиаф":
            addOptions(raceArchetype, ["Стандартный"]);
            break;
        case "Кенку":
            addOptions(raceArchetype, ["Стандартный"]);
            break;
        case "Совлин":
            addOptions(raceArchetype, ["Стандартный"]);
            break;
        case "Драконорожденный":
            addOptions(raceArchetype, ["Стандартный","Драконокровный","Равенит",
                "Металлический","Самоцветный","Цветной"]);
            break;
        default:
            break;
    }
}

