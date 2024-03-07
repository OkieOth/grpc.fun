<%
    import re
    import yacg.model.model as model
    import yacg.templateHelper as templateHelper
    import yacg.model.modelFuncs as modelFuncs
    import yacg.util.stringUtils as stringUtils

    templateFile = 'golang.mako'
    templateVersion = '1.0.1'

    packageName = templateParameters.get('modelPackage','<<PLEASE SET modelPackage TEMPLATE PARAM>>')


    def printGolangType(typeObj, isArray, isRequired, arrayDimensions = 1):
        ret = ''
        if typeObj is None:
            return '???'
        elif isinstance(typeObj, model.IntegerType):
            if typeObj.format is None or typeObj.format == model.IntegerTypeFormatEnum.INT32:
                ret = 'int32'
            else:
                ret = 'int'
        elif isinstance(typeObj, model.ObjectType):
            ret = 'interface{}'
        elif isinstance(typeObj, model.NumberType):
            if typeObj.format is None or typeObj.format == model.NumberTypeFormatEnum.DOUBLE:
                ret = 'float64'
            else:
                ret = 'float32'
        elif isinstance(typeObj, model.BooleanType):
            ret = 'bool'
        elif isinstance(typeObj, model.StringType):
            ret = 'string'
        elif isinstance(typeObj, model.BytesType):
            ret = 'byte'
        elif isinstance(typeObj, model.UuidType):
            ret = 'uuid.UUID'
        elif isinstance(typeObj, model.EnumType):
            ret = typeObj.name
        elif isinstance(typeObj, model.DateType):
            ret = 'time.Date'
        elif isinstance(typeObj, model.TimeType):
            ret = 'time.Time'
        elif isinstance(typeObj, model.DateTimeType):
            ret = 'time.Date'
        elif isinstance(typeObj, model.DictionaryType):
            ret = 'map[string]{}'.format(printGolangType(typeObj.valueType, False, True))
        elif isinstance(typeObj, model.ComplexType):
            ret = typeObj.name
        else:
            ret = '???'
        if isArray:
            ret = ("[]" * arrayDimensions) + ret
        if not isRequired:
            return "*" + ret
        else:
            return ret

    def getEnumDefaultValue(type):
        if type.default is not None:
            return secureEnumValues(type.default)
        else:
            return secureEnumValues(type.values[0])

    def secureEnumValues(value):
        pattern = re.compile("^[0-9]")
        return '_' + value if pattern.match(value) else value

    def isEnumDefaultValue(value, type):
        return getEnumDefaultValue(type) == secureEnumValues(value)


%>// Attention, this file is generated. Manual changes get lost with the next
// run of the code generation.
// created by yacg (template: ${templateFile} v${templateVersion})
package ${packageName}

import (
    "testing"
)

% for type in modelTypes:
    % if modelFuncs.isEnumType(type):
// how to test enums???
    % endif
% endfor

% for type in modelTypes:
    % if hasattr(type, "properties"):
func Test${type.name}(t *testing.T) {
    l := Make${type.name}()
        % for property in type.properties:
            % if property.isArray:
    if len(l.${stringUtils.toUpperCamelCase(property.name)}) != 0 {
        t.Errorf("Array field '${stringUtils.toUpperCamelCase(property.name)}' hasn't a 0 length, after creation")
    }

            % elif not property.required:
    if l.${stringUtils.toUpperCamelCase(property.name)}.IsSet {
        t.Errorf("Optional field '${stringUtils.toUpperCamelCase(property.name)}' is set, after creation, but should be unset")
    }

            % endif
        % endfor
}

    % endif
% endfor