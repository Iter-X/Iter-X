class ApiModel {
  bool? selected;
  String? title;
  String? baseUrl;

  ApiModel({this.selected, this.title, this.baseUrl});

  ApiModel.fromJson(Map<String, dynamic> json) {
    if (json["selected"] is bool) selected = json["selected"];
    if (json["title"] is String) title = json["title"];
    if (json["baseUrl"] is String) baseUrl = json["baseUrl"];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = <String, dynamic>{};
    data["selected"] = selected;
    data["title"] = title;
    data["baseUrl"] = baseUrl;
    return data;
  }
}
