package models

//input struct
type CreateEndpoint struct {
	Name_folderroute       string `json:"name_folderroute" valid:"required~name_folderroute: obligatorio"`
	Name_file              string `json:"name_file" valid:"required~name_file: obligatorio"`
	Name_model             string `json:"name_model" valid:"required~name_model: obligatorio"`
	Name_table             string `json:"name_table" valid:"required~name_table: obligatorio"`
	Name_camelcase         string `json:"name_camelcase" valid:"required~name_camelcase: obligatorio"`
	Name_lowercase         string `json:"name_lowercase" valid:"required~name_lowercase: obligatorio"`
	Name_loweruppercase    string `json:"name_loweruppercase" valid:"required~name_loweruppercase: obligatorio"`
	Name_abbreviationupper string `json:"name_abbreviationupper" valid:"required~name_abbreviationupper: obligatorio"`
	Name_abbreviationlower string `json:"name_abbreviationlower" valid:"required~name_abbreviationlower: obligatorio"`
	Name_message           string `json:"name_message" valid:"required~name_message: obligatorio"`
}
