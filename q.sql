SELECT
  CAST(REGEXP_EXTRACT(__key__.path, '^\"User\", ([0-9]+),') AS STRING) AS UserID,
  CAST(__key__.id AS STRING) AS CircleExhibitInfoID,
  CAST(EventExhibitCourseID AS STRING) AS EventExhibitCourseID,
  * EXCEPT(CircleExhibitInfoID,EventExhibitCourseID,__key__, __error__ , __has_error__)
FROM `tbf-tokyo.datastore_imports.CheckedCircleExhibit`